package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	name := "job_queue"
	addr := "amqp://guest:guest@localhost:5672/"
	queue := New(name, addr)
	message := []byte("message")
	// Attempt to push a message every 2 seconds
	for {
		time.Sleep(time.Second * 3)
		if err := queue.Push(message); err != nil {
			fmt.Printf("Push failed: %s\n", err)
		} else {
			fmt.Println("Push succeeded!")
		}
	}
}

type Session struct {
	name            string
	logger          *log.Logger
	connection      *amqp.Connection
	channel         *amqp.Channel
	done            chan bool
	notifyConnClose chan *amqp.Error
	notifyChanClose chan *amqp.Error
	notifyConfirm   chan amqp.Confirmation
	gotowe          bool
}

const (
	reconnectDelay = 5 * time.Second
	reInitDelay    = 2 * time.Second
	resendDelay    = 5 * time.Second
)

var (
	errNotConnected  = errors.New("not connected to a server")
	errAlreadyClosed = errors.New("already closed: not connected to the server")
	errShutdown      = errors.New("session is shutting down")
)

// New creates a new consumer state instance, and automatically
// attempts to connect to the server.
func New(name string, addr string) *Session {
	session := Session{
		logger: log.New(os.Stdout, "", log.LstdFlags),
		name:   name,
		done:   make(chan bool),
	}
	go session.pilnujPolaczenia(addr)
	return &session
}

// pilnujPolaczenia will wait for a connection error on
// notifyConnClose, and then continuously attempt to reconnect.
func (session *Session) pilnujPolaczenia(addr string) {
REDIAL:
	session.gotowe = false
	log.Println("Attempting to connect")

	conn, err := amqp.Dial(addr)

	if err != nil {
		log.Println("Failed to connect. Retrying...")
		select {
		case <-session.done:
			return
		case <-time.After(reconnectDelay):
		}
		goto REDIAL
	}

	session.connection = conn
	session.notifyConnClose = session.connection.NotifyClose(make(chan *amqp.Error))
	log.Println("Connected!")

REINIT:
	session.gotowe = false

	ch, err := conn.Channel()
	if err != nil {
		log.Println("Failed to initialize channel. Retrying...")

		select {
		case <-time.After(reInitDelay):
			goto REINIT
		case <-session.done:
			return
		}
	}

	err = session.przygotujQuełełe(ch)
	if err != nil {
		log.Println("Failed to initialize channel. Retrying...")

		select {
		case <-time.After(reInitDelay):
			goto REINIT
		case <-session.done:
			return
		}
	}

	session.channel = ch
	session.notifyChanClose = ch.NotifyClose(make(chan *amqp.Error))
	session.gotowe = true
	log.Println("Setup!")

	// czeka na koniec (sesja-Close, conn-Close, chan-Close)
	select {
	case <-session.notifyConnClose:
		log.Println("Connection closed. Reconnecting...")
		goto REDIAL
	case <-session.notifyChanClose:
		log.Println("Channel closed. Re-running init...")
		goto REINIT
	case <-session.done:
		return
	}
}

func (session *Session) przygotujQuełełe(ch *amqp.Channel) error {
	err := ch.Confirm(false)

	if err != nil {
		return err
	}
	_, err = ch.QueueDeclare(
		session.name,
		false, // Durable
		false, // Delete when unused
		false, // Exclusive
		false, // No-wait
		nil,   // Arguments
	)

	if err != nil {
		return err
	}

	session.notifyConfirm = ch.NotifyPublish(make(chan amqp.Confirmation, 1))

	return nil
}

// Push will push data onto the queue, and wait for a confirm.
// If no confirms are received until within the resendTimeout,
// it continuously re-sends messages until a confirm is received.
// This will block until the server sends a confirm. Errors are
// only returned if the push action itself fails, see UnsafePush.
func (session *Session) Push(data []byte) error {
	if !session.gotowe {
		return errors.New("failed to push push: not connected")
	}
	for {
		err := session.UnsafePush(data)
		if err != nil {
			session.logger.Println("Push failed. Retrying...")
			select {
			case <-session.done:
				return errShutdown
			case <-time.After(resendDelay):
			}
			continue
		}
		select {
		case confirm := <-session.notifyConfirm:
			if confirm.Ack {
				session.logger.Println("Push confirmed!")
				return nil
			}
		case <-time.After(resendDelay):
		}
		session.logger.Println("Push didn't confirm. Retrying...")
	}
}

// UnsafePush will push to the queue without checking for
// confirmation. It returns an error if it fails to connect.
// No guarantees are provided for whether the server will
// recieve the message.
func (session *Session) UnsafePush(data []byte) error {
	if !session.gotowe {
		return errNotConnected
	}
	return session.channel.Publish(
		"",           // Exchange
		session.name, // Routing key
		false,        // Mandatory
		false,        // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		},
	)
}

// Stream will continuously put queue items on the channel.
// It is required to call delivery.Ack when it has been
// successfully processed, or delivery.Nack when it fails.
// Ignoring this will cause data to build up on the server.
func (session *Session) Stream() (<-chan amqp.Delivery, error) {
	if !session.gotowe {
		return nil, errNotConnected
	}
	return session.channel.Consume(
		session.name,
		"",    // Consumer
		false, // Auto-Ack
		false, // Exclusive
		false, // No-local
		false, // No-Wait
		nil,   // Args
	)
}

// Close will cleanly shutdown the channel and connection.
func (session *Session) Close() error {
	if !session.gotowe {
		return errAlreadyClosed
	}
	err := session.channel.Close()
	if err != nil {
		return err
	}
	err = session.connection.Close()
	if err != nil {
		return err
	}
	close(session.done)
	session.gotowe = false
	return nil
}