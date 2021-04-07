package v2

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

var rozneQue = map[string]func(*Que, *amqp.Channel) error{
	"stdque": stdQue,
	//"szybki": przygotujSzybki,
	//"pewny": przygotujPewny,
}

// --- stdque ---

func stdQue(que *Que, chann *amqp.Channel) error {

	if _, err := chann.QueueDeclare(
		que.nazwa,
		true, // Durable
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

	que.konsumuj = stdKonsumuj
	return nil 
}

func stdKonsumuj(que *Que) error {
	wiadomosci, err := que.sesja.chann.Consume(
		que.nazwa, // name
		"",        // consumerTag, dostanę od serwera
		false,     // noAck
		false,     // exclusive
		false,     // noLocal - not suported by RabbitMQ
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		return fmt.Errorf("[Królik.Que] błąd Consume(): %v", err)
	}

	go func(wiadomosci <-chan amqp.Delivery) {
		for wiad := range wiadomosci {

			err := que.handler(wiad.Body)

			if err != nil {
				if err := wiad.Nack( /*multiple*/ false /*requeue*/, false); err != nil {
					log.Printf("błąd Nack: %v", err)
				}
			} else {
				if err := wiad.Ack( /*multiple*/ false); err != nil {
					log.Printf("błąd Nack: %v", err)
				}
			}
		
			// select {
			// case <-koniec:
			// 	break
			// default:
			// 	continue
			// }
		}
		log.Printf("[Królik.Que] Pobieranie zamknięte.")
		//que.sesja.Close()
		log.Printf("[Królik.Que] Pobieranie zamknięte chann.Close().")
	}(wiadomosci)

	return nil
}
