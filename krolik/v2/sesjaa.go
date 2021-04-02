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
		url:        url,
		done:       make(chan struct{}),
		przygotuj:  przygotuj,
		pilnujDone: make(chan bool),
	}
	rdy := make(chan bool)
	wgDone.Add(1)
	go sesja.pilnujSesji(rdy)
	<-rdy
	return sesja
}

func (sesja *sesjaa) Close() {
	sesja.czyOK = false
	close(sesja.done)
	<-sesja.pilnujDone
}

type funcPrzygotuj func(*amqp.Channel) error

type sesjaa struct {
	url              string
	czyOK            bool
	done             chan struct{}
	conn             *amqp.Connection
	chann            *amqp.Channel
	notifyConnClose  chan *amqp.Error
	notifyChannClose chan *amqp.Error
	przygotuj        funcPrzygotuj
	pilnujDone       chan bool
}

const (
	ponawianiePolaczPauza    = 7 * time.Second
	ponawianieKanalPauza     = 5 * time.Second
	ponawianiePrzygotujPauza = 10 * time.Second
)

func (sesja *sesjaa) pilnujSesji(rdy chan bool) {
	raz := sync.Once{}
	var chann *amqp.Channel

REDIAL:
	sesja.czyOK = false
	log.Printf("[Królik.Sesja] Łączę -> %q", sesja.url)

	conn, err := amqp.Dial(sesja.url)

	if err != nil {
		log.Printf("[Królik.Sesja] Błąd amqp.Dial(): %v", err)
		log.Printf("[Królik.Sesja] Ponawiam za %v...", ponawianiePolaczPauza)

		select {
		case <-time.After(ponawianiePolaczPauza):
			goto REDIAL
		case <-sesja.done:
			goto DONE
		}
	}

	sesja.conn = conn
	sesja.notifyConnClose = sesja.conn.NotifyClose(make(chan *amqp.Error))
	log.Println("[Królik.Sesja] Połączony.")

REINIT:
	sesja.czyOK = false

	chann, err = conn.Channel()
	if err != nil {
		log.Printf("[Królik.Sesja] Błąd conn.Channel(): %v", err)
		log.Printf("[Królik.Sesja] Ponawiam za %v...", ponawianieKanalPauza)

		select {
		case <-time.After(ponawianieKanalPauza):
			goto REINIT
		case err := <-sesja.notifyConnClose:
			log.Printf("[Królik.Sesja] Połączenie zamknięte: %v", err)
			goto REDIAL
		case <-sesja.done:
			goto DONE
		}
	}

	err = sesja.przygotuj(chann)
	if err != nil {
		log.Printf("[Królik.Sesja] Błąd sesja.przygotuj(): %v", err)
		log.Printf("[Królik.Sesja] Ponawiam za %v...", ponawianiePrzygotujPauza)

		select {
		case <-time.After(ponawianiePrzygotujPauza):
			goto REINIT
		case err := <-sesja.notifyConnClose:
			log.Printf("[Królik.Sesja] Połączenie zamknięte: %v", err)
			goto REDIAL
		case <-sesja.done:
			goto DONE
		}
	}

	sesja.chann = chann
	sesja.notifyChannClose = chann.NotifyClose(make(chan *amqp.Error))

	// Sesja otwarta.
	sesja.czyOK = true
	log.Println("[Królik.Sesja] Rdy.")
	raz.Do(func() { rdy <- true })

	// Czekamy na jakiś koniec.
	select {
	case err := <-sesja.notifyConnClose:
		log.Printf("[Królik.Sesja] Połączenie zamknięte: %v", err)
		goto REDIAL
	case err := <-sesja.notifyChannClose:
		log.Printf("[Królik.Sesja] Kanał zamknięty: %v", err)
		goto REINIT
	case <-sesja.done:
		goto DONE
	}

DONE:
	log.Println("[Królik.Sesja] Done.")
	if sesja.chann != nil {
		err = sesja.chann.Close()
		if err != nil {
			log.Printf("[Królik.Sesja] błąd sesja.chan.Close(): %q", err)
		}
	}
	if sesja.conn != nil {
		err = sesja.conn.Close()
		if err != nil {
			log.Printf("[Królik.Sesja] błąd sesja.conn.Close(): %q", err)
		}
	}
	sesja.pilnujDone <- true
}
