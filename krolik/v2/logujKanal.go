package v2

import (
	"log"

	"github.com/streadway/amqp"
)

func logujKanal(ch *amqp.Channel) {
	go func() {
		for x := range ch.NotifyCancel(make(chan string)) {
			log.Println("NotifyCancel:", x)
		}
		log.Println("NotifyCancel closed")
	}()
	go func() {
		for x := range ch.NotifyClose(make(chan *amqp.Error)) {
			log.Println("NotifyClose:", x)
		}
		log.Println("NotifyClose closed")
	}()
	go func() {
		// Before RabbitMQ 3.2, the RabbitMQ team deprecated the use of Channel.Flow, 
		// replacing it with a mechanism called TCP Backpressure
		for x := range ch.NotifyFlow(make(chan bool)) {
			log.Println("NotifyFlow", x)
		}
		log.Println("NotifyFlow closed")
	}()
	go func() {
		for x := range ch.NotifyPublish(make(chan amqp.Confirmation)) {
			log.Println("NotifyPublish", x)
		}
		log.Println("NotifyPublish closed")
	}()
	go func() {
		for x := range ch.NotifyReturn(make(chan amqp.Return)) {
			log.Println("NotifyReturn", x)
		}
		log.Println("NotifyReturn closed")
	}()
}
