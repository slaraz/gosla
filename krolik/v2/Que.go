package v2

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

type QueHandler func([]byte) error

type Que struct {
	sesja   *sesjaa
	nazwa   string
	bindEx  string
	handler QueHandler
	stopWrk chan bool
	wrkWG   *sync.WaitGroup
}

func MusiQuełełe(url, nazwa, bindEx, rodzaj string, handler QueHandler) *Que {
	que := &Que{
		nazwa:   nazwa,
		bindEx:  bindEx,
		handler: handler,
		stopWrk: make(chan bool),
		wrkWG:   &sync.WaitGroup{},
	}

	wybranyRodzajQue, ok := rozneQue[rodzaj]
	if !ok {
		log.Fatalln("nieznany rodzaj quełełe")
	}

	przygotuj := func(chann *amqp.Channel, log *log.Logger) error {
		log.Printf("Przygotowuję [%s->%q]", rodzaj, bindEx)
		qp := queParam{nazwa, bindEx}
		err := wybranyRodzajQue.przygotuj(qp, chann)
		if err != nil {
			return err
		}
		wiadomosci, err := wybranyRodzajQue.konsumuj(qp, chann)
		if err != nil {
			return err
		}
		que.podepnijHandler(wiadomosci, wybranyRodzajQue.odbierz)
		return nil
	}

	nazwaSesji := fmt.Sprintf("QUE<%s>", nazwa)
	sesja := otworz(url, przygotuj, nazwaSesji)
	que.sesja = sesja

	return que
}

func (que *Que) Close() {
	// Zamykamy wątki odbierające.
	close(que.stopWrk)
	// Czekamy aż się zakończą.
	que.wrkWG.Wait()
	// Zamykamy połączenie z Rabbitem.
	if que.sesja != nil {
		que.sesja.close()
	}
}

func (que *Que) podepnijHandler(wiadomosci <-chan amqp.Delivery, odbierzDelivery func(amqp.Delivery, QueHandler) error) {
	que.stopWrk = make(chan bool)
	que.wrkWG.Add(1)
	go func() {
		log := log.New(os.Stdout, "[Que.Wrk] ", 0)
	petla:
		for {
			select {
			case wiad, ok := <-wiadomosci:
				if !ok {
					log.Print("kanał wiadomosci zamknięty")
					break petla
				}
				// Obsługa wiadomości przez klienta.
				err := odbierzDelivery(wiad, que.handler)
				if err != nil {
					log.Printf("błąd odbierzDelivery(): %v", err)
				}

			case <-time.After(3 * time.Second):
				log.Printf("brak wiadomości...")

			case <-que.stopWrk:
				log.Print("zatrzymuję")
				break petla
			}
		}
		que.wrkWG.Done()
		log.Printf("DONE")
	}()
}
