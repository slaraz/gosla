package v2

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func MusiNowyExchanger(conf string) *Ex {
	konfEx, err := parsujKonf(conf)
	if err != nil {
		log.Fatalf("błąd parsujYAML(): %v", err)
	}
	log.Println(konfEx)

	ex, err := nowyEx(konfEx)
	if err != nil {
		log.Fatalf("błąd NowyExchanger(): %v", err)
	}

	return ex
}

func (ex *Ex) PublikujJSON(v interface{}) error {
	bajty, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("błąd json.Marshal(): %v", err)
	}
	err = ex.publikuj(bajty)
	if err != nil {
		return fmt.Errorf("błąd ex.publikuj(): %v", err)
	}
	return nil
}

func (ex *Ex) Close() error {
	if ex.ch != nil {
		if err := ex.ch.Close(); err != nil {
			return fmt.Errorf("błąd ch.Close(): %v", err)
		}
	}
	return nil
}

type Ex struct {
	ch         *amqp.Channel
	nazwa      string
	kind       string
	publikuj   func(dane []byte) error
	routingKey string
}

type konfiguracjaEx struct {
	url, nazwa, kind, rodzaj, routingKey string
}

func nowyEx(konfEx konfiguracjaEx) (*Ex, error) {
	ex := &Ex{
		nazwa:      konfEx.nazwa,
		kind:       konfEx.kind,
		routingKey: konfEx.routingKey,
	}

	ch, err := kanal(konfEx.url)
	if err != nil {
		return nil, fmt.Errorf("błąd ex.polacz(): %v", err)
	}
	ex.ch = ch

	if przygotujEx, ok := rozneEx[konfEx.rodzaj]; ok {
		if err := przygotujEx(ex); err != nil {
			return nil, fmt.Errorf("błąd przygotujEx(): %v", err)
		}
	} else {
		return nil, fmt.Errorf("błąd rodzajeEx[]: nieznany rodzaj exchangera")
	}

	return ex, nil
}
