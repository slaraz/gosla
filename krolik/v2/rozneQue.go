package v2

import (
	"fmt"

	"github.com/streadway/amqp"
)

type queParam struct {
	nazwa  string
	bindTo string
}

type rodzajQue struct {
	przygotuj func(queParam, *amqp.Channel) error
	konsumuj  func(queParam, *amqp.Channel) (<-chan amqp.Delivery, error)
	odbierz   func(amqp.Delivery, QueHandler) error
}

var rozneQue = map[string]rodzajQue{
	"stdque": {stdQue, stdKonsumuj, stdOdbierz},
	//"szybki": przygotujSzybki,
	//"pewny": przygotujPewny,
}

// --- stdque ---

func stdQue(que queParam, chann *amqp.Channel) error {
	if _, err := chann.QueueDeclare(
		que.nazwa,
		true,  // Durable
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
		que.bindTo, // sourceExchange
		false,      // noWait
		nil,        // arguments
	); err != nil {
		return fmt.Errorf("chann.QueueBind(): %v", err)
	}
	return nil
}

func stdKonsumuj(que queParam, chann *amqp.Channel) (<-chan amqp.Delivery, error) {
	return chann.Consume(
		que.nazwa, // name
		"",        // consumerTag, dostanę od serwera
		false,     // noAck
		false,     // exclusive
		false,     // noLocal - not suported by RabbitMQ
		false,     // noWait
		nil,       // arguments
	)
}

func stdOdbierz(wiad amqp.Delivery, handler QueHandler) error {
	err := handler(wiad.Body)
	if err != nil {
		if err := wiad.Nack( /*multiple*/ false /*requeue*/, false); err != nil {
			return fmt.Errorf("wiad.Nack(): %v", err)
		}
	} else {
		if err := wiad.Ack( /*multiple*/ false); err != nil {
			return fmt.Errorf("wiad.Ack(): %v", err)
		}
	}
	return nil
}
