package v2

import (
	"log"

	"github.com/streadway/amqp"
)

type Que struct {
	sesja    *sesjaa
	nazwa    string
	bindEx   string
	konsumuj func(que *Que) error
	handler func([]byte) error
}

func MusiQuełełe(url, nazwa, bindEx, rodzaj string, handler func([]byte) error) *Que {
	que := &Que{
		nazwa:  nazwa,
		bindEx: bindEx,
		handler: handler,
	}

	wybranyRodzajQue, ok := rozneQue[rodzaj]
	if !ok {
		log.Fatalln("rozneEx[]: nieznany rodzaj quełełe")
	}

	przygotuj := func(chann *amqp.Channel) error {
		log.Printf("[Królik.Que] Przygotowuję -> [%q:%s:%s]", nazwa, bindEx, rodzaj)
		return wybranyRodzajQue(que, chann)
	}

	sesja := Otworz(url, przygotuj)
	que.sesja = sesja

	que.konsumuj(que)

	return que
}

func (que *Que) Close() {
	if que.sesja != nil {
		que.sesja.Close()
	}
}
