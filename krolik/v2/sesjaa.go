package v2

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

const ponawianiePauza = 5 * time.Second

type funcPrzygotuj func(*amqp.Channel, *log.Logger) error

type sesjaa struct {
	nazwa     string
	czyGotowa bool
	chann     *amqp.Channel
	Chann     chan *amqp.Channel
	doneRQ    chan struct{}
	doneACK   chan bool
}

func otworz(url string, przygotuj funcPrzygotuj, nazwaSesji string) *sesjaa {
	sesja := &sesjaa{
		nazwa:   nazwaSesji,
		doneRQ:  make(chan struct{}),
		doneACK: make(chan bool),
		Chann:   make(chan *amqp.Channel),
	}
	started := make(chan bool)
	go sesja.pilnujSesji(url, started, przygotuj)
	<-started
	return sesja
}

func (sesja *sesjaa) close() {
	sesja.czyGotowa = false
	close(sesja.doneRQ)
	<-sesja.doneACK
}

func (sesja *sesjaa) pilnujSesji(url string, started chan bool, przygotuj funcPrzygotuj) {
	log := log.New(os.Stdout, fmt.Sprintf("[%s] ", sesja.nazwa), 0)
	raz := sync.Once{}
	var chann *amqp.Channel
	var notifyConnClose chan *amqp.Error
	var notifyChannClose chan *amqp.Error

REDIAL:
	sesja.czyGotowa = false
	log.Printf("Łączę -> %q", url)

	conn, err := amqp.Dial(url)

	if err != nil {
		log.Printf("Błąd amqp.Dial(): %v", err)
		log.Printf("Ponawiam za %v...", ponawianiePauza)

		select {
		case <-time.After(ponawianiePauza):
			goto REDIAL
		case <-sesja.doneRQ:
			goto KONIEC
		}
	}

	notifyConnClose = conn.NotifyClose(make(chan *amqp.Error))
	//log.Printf("Połączony <=> AMQP %d-%d", conn.Major, conn.Minor)

REINIT:
	sesja.czyGotowa = false

	chann, err = conn.Channel()
	if err != nil {
		log.Printf("Błąd conn.Channel(): %v", err)
		log.Printf("Ponawiam za %v...", ponawianiePauza)

		select {
		case <-time.After(ponawianiePauza):
			goto REINIT
		case err := <-notifyConnClose:
			log.Printf("Połączenie zamknięte: %v", err)
			goto REDIAL
		case <-sesja.doneRQ:
			goto KONIEC
		}
	}

	// ---

	err = przygotuj(chann, log)
	if err != nil {
		log.Printf("Błąd sesja.przygotuj(): %v", err)
		log.Printf("Ponawiam za %v...", ponawianiePauza)

		select {
		case <-time.After(ponawianiePauza):
			goto REINIT
		case err := <-notifyConnClose:
			log.Printf("Połączenie zamknięte: %v", err)
			goto REDIAL
		case <-sesja.doneRQ:
			goto KONIEC
		}
	}

	sesja.chann = chann
	notifyChannClose = chann.NotifyClose(make(chan *amqp.Error))
	sesja.czyGotowa = true
	log.Printf("READY")
	raz.Do(func() { started <- true })

	// Sesja pracuje.
pracuje:
	// Czekamy na jakiś koniec.
	select {
	case sesja.Chann <- sesja.chann:
		goto pracuje
	case err := <-notifyConnClose:
		<-notifyChannClose // HACK: tej linijki szukałem 1/2 dnia, zamyka (chan *amqp.Delivery)
		log.Printf("Połączenie zamknięte: %v", err)
		goto REDIAL
	case err := <-notifyChannClose:
		log.Printf("Kanał zamknięty: %v", err)
		goto REINIT
	case <-sesja.doneRQ:
		goto KONIEC
	}

KONIEC:
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
	log.Printf("KONIEC.")
}
