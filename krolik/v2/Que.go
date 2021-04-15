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
	konsumKoniec  chan struct{}
	wrkWG         *sync.WaitGroup
	wybranyRodzaj rodzajQue
	sesja         *sesjaa
	Pusta         chan struct{}
}

func MusiQuełełe(url string, qp queParam, rodzaj string) *Que {
	wybranyRodzajQue, ok := rozneQue[rodzaj]
	if !ok {
		log.Fatalln("nieznany rodzaj quełełe")
	}
	que := &Que{
		qp:            qp,
		rodzaj:        rodzaj,
		wrkWG:         &sync.WaitGroup{},
		wybranyRodzaj: wybranyRodzajQue,
		Pusta:         make(chan struct{}),
	}
	que.sesja = otworz(url, que.przygotuj, fmt.Sprintf("QUE %s", qp.nazwa))
	return que
}

func (que *Que) przygotuj(chann *amqp.Channel, log *log.Logger) error {
	log.Printf("Przygotowuję [%s bind->%q]", que.rodzaj, que.qp.bindTo)
	q, err := que.wybranyRodzaj.przygotuj(que.qp, chann)
	if err != nil {
		return err
	}
	log.Printf("konsumentów: %d, wiadomości: %d", q.Consumers, q.Messages)
	return nil
}

func (que *Que) Close() {
	// Zamykamy wątki odbierające.
	if que.konsumKoniec != nil {
		close(que.konsumKoniec)
	}
	// Czekamy aż się zakończą.
	que.wrkWG.Wait()
	// Zamykamy połączenie z Rabbitem.
	if que.sesja != nil {
		que.sesja.close()
	}
}

func (que *Que) Konsumuj(handler QueHandler) {
	que.konsumKoniec = make(chan struct{})
	que.wrkWG.Add(1)
	go func() {
		defer que.wrkWG.Done()
		log := log.New(os.Stdout, "[Que.Wrk] ", 0)
		var chann *amqp.Channel
	INIT:
		log.Print("Czekam na połączenie...")
		select {
		case chann = <-que.sesja.Chann:
			break
		case <-que.konsumKoniec:
			log.Printf("KONIEC.")
			return
		}
		wiadChann, err := que.wybranyRodzaj.konsumuj(que.qp, chann)
		if err != nil {
			log.Printf("błąd konsumuj(): %v", err)
			goto INIT
		}
		log.Print("KONSUMUJĘ")
		for {
			select {
			case wiad, ok := <-wiadChann:
				if !ok {
					log.Print("kanał wiadomosci zamknięty")
					goto INIT
				}
				// Obsługa wiadomości przez klienta.
				err := que.wybranyRodzaj.odbierz(wiad, handler)
				if err != nil {
					log.Printf("błąd odbierzDelivery(): %v", err)
				}

			//case <-time.After(3 * time.Second):

			case <-time.After(time.Second):
				// Jeśli przez sekundę nic nie przyszło to kolejka jest pusta.
				select {
				case que.Pusta <- struct{}{}:
					break
				default:
					break
				}

			case <-que.konsumKoniec:
				log.Printf("KONIEC.")
				return

			}
		}
	}()
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

func (que *Que) UsunWiadomosci() error {
	if !que.sesja.czyGotowa {
		return fmt.Errorf("brak połączenia")
	}
	ile, err := que.sesja.chann.QueuePurge(que.qp.nazwa, true)
	if err != nil {
		return fmt.Errorf("chann.QueuePurge(): %v", err)
	}
	if ile > 0 {
		log.Printf("usunąłem %d wiad", ile)
		return nil
	}
	return nil
}
