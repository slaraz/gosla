package v2

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func MusiQuełełe(url, nazwa, bindEx, rodzaj string, handler func([]byte) error) *Que {
	que := &Que{
		nazwa:  nazwa,
		bindEx: bindEx,
		handler: handler,
	}

	przygotuj := func(chann *amqp.Channel) error {
		log.Printf("[Królik.Que] Przygotowuję -> [%q:%s:%s]", nazwa, bindEx, rodzaj)
		return przygotujQue(que, chann)
	}

	sesja := Otworz(url, przygotuj)
	que.sesja = sesja

	return que
}

func (que *Que) konsumuj(handler func([]byte) error) {
	
	log.Printf("[RabbitMQ.Cons] <~~ [%s]", que.nazwa)

}

type Que struct {
	sesja  *sesjaa
	nazwa  string
	bindEx string
	handler func([]byte) error
}

func (que *Que) OdbierzJSON(v interface{}) error {
	return nil
}

func (que *Que) Close() {
	if que.sesja != nil {
		que.sesja.Close()
	}
}

func przygotujQue(que *Que, chann *amqp.Channel) error {

	if _, err := chann.QueueDeclare(
		que.nazwa,
		false, // Durable
		false, // Delete when unused
		false, // Exclusive
		false, // No-wait
		nil,   // Arguments
	); err != nil {
		return fmt.Errorf("chann.QueueDeclare(): %v", err)
	}

	if err := chann.QueueBind(
		que.nazwa,  // name of the queue
		"",         // bindingKey
		que.bindEx, // sourceExchange
		false,      // noWait
		nil,        // arguments
	); err != nil {
		return fmt.Errorf("chann.QueueBind(): %v", err)
	}

	wiadomosci, err := chann.Consume(
		que.nazwa, // name
		"",        // consumerTag, dostanę od serwera
		false,     // noAck
		false,     // exclusive
		false,     // noLocal - not suported by RabbitMQ
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("[Królik.Que] błąd Consume(): %v", err)
	}

	go func(wiadomosci <-chan amqp.Delivery) {
		for wiad := range wiadomosci {
			err := que.handler(wiad.Body)
			if err != nil {
				wiad.Nack( /*multiple*/ false /*requeue*/, false)
			} else {
				wiad.Ack( /*multiple*/ false)
			}
		}
		log.Printf("[Królik.Que] Pobieranie zamknięte. !!!!!!!!!!!!!!!!!!!!!!!!!!!")
		chann.Close()
	}(wiadomosci)

	return nil
}
