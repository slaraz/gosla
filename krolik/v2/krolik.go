package v2

import (
	"fmt"

	"github.com/streadway/amqp"
)

var kr *krolik2

type krolik2 struct {
	conn map[string]*amqp.Connection
	exs  map[string]*Ex
	//ques map[string]*Que
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
