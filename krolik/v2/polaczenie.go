package v2

import (
	"fmt"

	"github.com/streadway/amqp"
)

var jedenKrolikByRzadzicWszystkimi *krolik2

type krolik2 struct {
	conn map[string]*amqp.Connection
}

func kanal(url string) (*amqp.Channel, error) {
	if jedenKrolikByRzadzicWszystkimi == nil {
		jedenKrolikByRzadzicWszystkimi = &krolik2{
			conn: map[string]*amqp.Connection{},
		}
	}

	conn, err := jedenKrolikByRzadzicWszystkimi.polaczenie(url)
	if err != nil {
		return nil, fmt.Errorf("błąd kr.polaczenie(): %v", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("błąd conn.Channel(): %v", err)
	}
	return ch, nil
}

func (kr *krolik2) polaczenie(url string) (*amqp.Connection, error) {
	var conn *amqp.Connection
	conn, ok := kr.conn[url]
	if !ok {
		newConn, err := amqp.Dial(url)
		if err != nil {
			return nil, fmt.Errorf("błąd amqp.Dial(): %v", err)
		}
		kr.conn[url] = newConn
		conn = newConn
	}
	return conn, nil
}

