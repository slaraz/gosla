package v2

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

const (
	ponawianiePolaczPauza    = 5 * time.Second
	ponawianieKanalPauza     = 5 * time.Second
	ponawianiePrzygotujPauza = 5 * time.Second
)

type funcPrzygotuj func(*amqp.Channel, *log.Logger) error

type sesjaa struct {
	nazwa   string
	czyOK   bool
	doneRQ  chan struct{}
	chann   *amqp.Channel
	doneACK chan bool
}

func otworz(url string, przygotuj funcPrzygotuj, nazwaSesji string) *sesjaa {
	wgDone := &sync.WaitGroup{}
	sesja := &sesjaa{
		nazwa:   nazwaSesji,
		doneRQ:  make(chan struct{}),
		doneACK: make(chan bool),
	}
	rdy := make(chan bool)
	wgDone.Add(1)
	go sesja.pilnujSesji(url, rdy, przygotuj)
	<-rdy
	return sesja
}

func (sesja *sesjaa) close() {
	sesja.czyOK = false
	close(sesja.doneRQ)
	<-sesja.doneACK
	log.Printf("%q doneACK", sesja.nazwa)
}

func (sesja *sesjaa) pilnujSesji(url string, rdy chan bool, przygotuj funcPrzygotuj) {
	log := log.New(os.Stdout, fmt.Sprintf("[%s] ", sesja.nazwa), 0)
	raz := sync.Once{}
	var chann *amqp.Channel
	var notifyConnClose chan *amqp.Error
	var notifyChannClose chan *amqp.Error

REDIAL:
	sesja.czyOK = false
	log.Printf("Łączę -> %q", url)

	conn, err := amqp.Dial(url)

	if err != nil {
		log.Printf("Błąd amqp.Dial(): %v", err)
		log.Printf("Ponawiam za %v...", ponawianiePolaczPauza)

		select {
		case <-time.After(ponawianiePolaczPauza):
			goto REDIAL
		case <-sesja.doneRQ:
			goto DONE
		}
	}

	notifyConnClose = conn.NotifyClose(make(chan *amqp.Error))
	//log.Printf("Połączony <=> AMQP %d-%d", conn.Major, conn.Minor)

REINIT:
	sesja.czyOK = false

	chann, err = conn.Channel()
	if err != nil {
		log.Printf("Błąd conn.Channel(): %v", err)
		log.Printf("Ponawiam za %v...", ponawianieKanalPauza)

		select {
		case <-time.After(ponawianieKanalPauza):
			goto REINIT
		case err := <-notifyConnClose:
			log.Printf("Połączenie zamknięte: %v", err)
			goto REDIAL
		case <-sesja.doneRQ:
			goto DONE
		}
	}

	// ---

	err = przygotuj(chann, log)
	if err != nil {
		log.Printf("Błąd sesja.przygotuj(): %v", err)
		log.Printf("Ponawiam za %v...", ponawianiePrzygotujPauza)

		select {
		case <-time.After(ponawianiePrzygotujPauza):
			goto REINIT
		case err := <-notifyConnClose:
			log.Printf("Połączenie zamknięte: %v", err)
			goto REDIAL
		case <-sesja.doneRQ:
			goto DONE
		}
	}

	sesja.chann = chann
	notifyChannClose = chann.NotifyClose(make(chan *amqp.Error))
	sesja.czyOK = true
	log.Printf("READY")
	raz.Do(func() { rdy <- true })

	// Sesja pracuje.

	// Czekamy na jakiś koniec.
	select {
	case err := <-notifyConnClose:
		<-notifyChannClose // HACK: tej linijki szukałem 1/2 dnia, zamyka (chan *amqp.Delivery)
		log.Printf("Połączenie zamknięte: %v", err)
		goto REDIAL
	case err := <-notifyChannClose:
		log.Printf("Kanał zamknięty: %v", err)
		goto REINIT
	case <-sesja.doneRQ:
		goto DONE
	}

DONE:
	log.Printf("zamykam")
	if conn != nil {
		//log.Printf("conn.Close()")
		err := conn.Close()
		//log.Printf("conn.Closed")
		if err != nil {
			log.Printf("błąd sesja.conn.Close(): %q", err)
		}
	}
	//log.Printf("pilnujDone <- true")
	sesja.doneACK <- true
	log.Printf("DONE")
}

var start time.Time
