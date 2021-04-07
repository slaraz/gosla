package v2

import (
	"log"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

func Otworz(url string, przygotuj funcPrzygotuj) *sesjaa {
	wgDone := &sync.WaitGroup{}
	sesja := &sesjaa{
		doneRQ:       make(chan struct{}),
		doneACK: make(chan bool),
	}
	rdy := make(chan bool)
	wgDone.Add(1)
	go sesja.pilnujSesji(url, rdy, przygotuj)
	<-rdy
	return sesja
}

func (sesja *sesjaa) Close() {
	sesja.czyOK = false
	close(sesja.doneRQ)
	<-sesja.doneACK
	log.Println("[Królik.Sesja] sesja closed")
}

type funcPrzygotuj func(*amqp.Channel) error

type sesjaa struct {
	czyOK            bool
	doneRQ             chan struct{}
	chann            *amqp.Channel
	doneACK       chan bool
}

const (
	ponawianiePolaczPauza    = 7 * time.Second
	ponawianieKanalPauza     = 5 * time.Second
	ponawianiePrzygotujPauza = 10 * time.Second
)

func (sesja *sesjaa) pilnujSesji(url string, rdy chan bool, przygotuj funcPrzygotuj) {
	raz := sync.Once{}
	var chann *amqp.Channel
	var notifyConnClose chan *amqp.Error
	var notifyChannClose chan *amqp.Error

REDIAL:
	sesja.czyOK = false
	log.Printf("[Królik.Sesja] Łączę -> %q", url)

	conn, err := amqp.Dial(url)

	if err != nil {
		log.Printf("[Królik.Sesja] Błąd amqp.Dial(): %v", err)
		log.Printf("[Królik.Sesja] Ponawiam za %v...", ponawianiePolaczPauza)

		select {
		case <-time.After(ponawianiePolaczPauza):
			goto REDIAL
		case <-sesja.doneRQ:
			goto DONE
		}
	}

	notifyConnClose = conn.NotifyClose(make(chan *amqp.Error))
	log.Printf("[Królik.Sesja] Połączony <=> AMQP %d-%d", conn.Major, conn.Minor)

REINIT:
	sesja.czyOK = false

	chann, err = conn.Channel()
	if err != nil {
		log.Printf("[Królik.Sesja] Błąd conn.Channel(): %v", err)
		log.Printf("[Królik.Sesja] Ponawiam za %v...", ponawianieKanalPauza)

		select {
		case <-time.After(ponawianieKanalPauza):
			goto REINIT
		case err := <-notifyConnClose:
			log.Printf("[Królik.Sesja] Połączenie zamknięte: %v", err)
			goto REDIAL
		case <-sesja.doneRQ:
			goto DONE
		}
	}

	// ---

	err = przygotuj(chann)
	if err != nil {
		log.Printf("[Królik.Sesja] Błąd sesja.przygotuj(): %v", err)
		log.Printf("[Królik.Sesja] Ponawiam za %v...", ponawianiePrzygotujPauza)

		select {
		case <-time.After(ponawianiePrzygotujPauza):
			goto REINIT
		case err := <-notifyConnClose:
			log.Printf("[Królik.Sesja] Połączenie zamknięte: %v", err)
			goto REDIAL
		case <-sesja.doneRQ:
			goto DONE
		}
	}

	sesja.chann = chann
	notifyChannClose = chann.NotifyClose(make(chan *amqp.Error))
	sesja.czyOK = true
	log.Println("[Królik.Sesja] Rdy.")
	raz.Do(func() { rdy <- true })

	// Sesja pracuje.

	// Czekamy na jakiś koniec.
	select {
	case err := <-notifyConnClose:
		<-notifyChannClose // HACK: tej linijki szukałem 1/2 dnia, zamyka (chan *amqp.Delivery)
		log.Printf("[Królik.Sesja] Połączenie zamknięte: %v", err)
		goto REDIAL
	case err := <-notifyChannClose:
		log.Printf("[Królik.Sesja] Kanał zamknięty: %v", err)
		goto REINIT
	case <-sesja.doneRQ:
		goto DONE
	}

DONE:
	if chann != nil {
		err := chann.Close()
		if err != nil {
			log.Printf("błąd sesja.chann.Close(): %v", err)
		}
	}
	//sesja.czekajNaWorkers(time.Second)

	log.Println("[Królik.Sesja] DONE.", time.Since(start))
	if conn != nil {
		err := conn.Close()
		log.Println("sesja.conn.Closed", time.Since(start))
		if err != nil {
			log.Printf("[Królik.Sesja] błąd sesja.conn.Close(): %q", err)
		}
	}
	log.Println("sesja.pilnujDone <- true", time.Since(start))
	sesja.doneACK <- true
}

var start time.Time

