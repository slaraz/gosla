package v2

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

type QueHandler func([]byte) error

type Que struct {
	qp            queParam
	rodzaj        string
	handler       QueHandler
	koniecKonsum  chan bool
	wrkWG         *sync.WaitGroup
	wybranyRodzaj rodzajQue
	sesja         *sesjaa
}

func MusiQuełełe(url string, qp queParam, rodzaj string) *Que {
	wybranyRodzajQue, ok := rozneQue[rodzaj]
	if !ok {
		log.Fatalln("nieznany rodzaj quełełe")
	}
	que := &Que{
		qp:            qp,
		rodzaj:        rodzaj,
		koniecKonsum:  make(chan bool),
		wrkWG:         &sync.WaitGroup{},
		wybranyRodzaj: wybranyRodzajQue,
	}
	que.sesja = otworz(url, que.przygotuj, fmt.Sprintf("QUE %s", qp.nazwa))
	return que
}

func (que *Que) Close() {
	// Zamykamy wątki odbierające.
	close(que.koniecKonsum)
	// Czekamy aż się zakończą.
	que.wrkWG.Wait()
	// Zamykamy połączenie z Rabbitem.
	if que.sesja != nil {
		que.sesja.close()
	}
}

func (que *Que) przygotuj(chann *amqp.Channel, log *log.Logger) error {
	log.Printf("Przygotowuję [%s bind->%q]", que.rodzaj, que.qp.bindTo)
	q, err := que.wybranyRodzaj.przygotuj(que.qp, chann)
	if err != nil {
		return err
	}
	log.Printf("konsumentów: %d, wiadomości: %d", q.Consumers, q.Messages)

	//
	// TODO: kiedy ponawiamy połączenie odtworzyć Workery Kosumujące.

	// if que.handler != nil {
	// 	if err := que.konsumuj(chann); err != nil {
	// 		return err
	// 	}
	// }

	return nil
}

func (que *Que) Konsumuj(handler QueHandler) error {
	que.handler = handler
	chann := que.sesja.chann
	wiadChann, err := que.wybranyRodzaj.konsumuj(que.qp, chann)
	if err != nil {
		return err
	}
	//que.podepnijHandler(wiadChann, wybranyRodzajQue.odbierz)
	// }
	// func (que *Que) podepnijHandler(wiadChann <-chan amqp.Delivery, odbierzDelivery func(amqp.Delivery, QueHandler) error) {
	que.koniecKonsum = make(chan bool)
	que.wrkWG.Add(1)
	go func() {
		log := log.New(os.Stdout, "[Que.Wrk] ", 0)
	petla:
		for {
			select {
			case wiad, ok := <-wiadChann:
				if !ok {
					log.Print("kanał wiadomosci zamknięty")
					break petla
				}
				// Obsługa wiadomości przez klienta.
				err := que.wybranyRodzaj.odbierz(wiad, que.handler)
				if err != nil {
					log.Printf("błąd odbierzDelivery(): %v", err)
				}

			case <-time.After(3 * time.Second):
				log.Printf("brak wiadomości...")

			case <-que.koniecKonsum:
				log.Print("zatrzymuję")
				break petla
			}
		}
		que.wrkWG.Done()
		log.Printf("KONIEC.")
	}()
	return nil
}

func (que *Que) GetJSON(v interface{}) error {
	if !que.sesja.czyGotowa {
		return fmt.Errorf("brak połączenia")
	}
	msg, ok, err := que.sesja.chann.Get(que.qp.nazwa, true)
	if err != nil {
		return fmt.Errorf("chann.Get(): %v", err)
	}
	if !ok {
		// Nie odebraliśmy żadnej wiadomości, kolejka była pusta.
		return nil
	}
	//ileWiad := msg.MessageCount
	err = json.Unmarshal(msg.Body, v)
	if err != nil {
		return fmt.Errorf("json.Unmarshal(): %v", err)
	}
	return nil
}
