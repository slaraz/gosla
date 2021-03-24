package v2

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
	"gopkg.in/yaml.v3"
)

func MusiNowyExchanger(conf string) *Ex {
	ex, err := nowyExchanger(conf)
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

// ---

type konfig struct {
	Url       string `yaml:"url"` // https://www.rabbitmq.com/uri-spec.html
	Exchanger struct {
		Nazwa      string `yaml:"nazwa"`
		Kind       string `yaml:"kind"`
		Rodzaj     string `yaml:"rodzaj"`
		RoutingKey string `yaml:"routingKey,omitempty"`
	} `yaml:"exchanger"`
}

func nowyExchanger(conf string) (*Ex, error) {
	konf := konfig{}
	err := yaml.Unmarshal([]byte(conf), &konf)
	if err != nil {
		fmt.Printf("błąd yaml.Unmarshal(): %v\n", err)
	}
	fmt.Println(konf)

	if kr == nil {
		kr = &krolik2{
			conn: map[string]*amqp.Connection{},
			exs:  map[string]*Ex{},
		}
	}

	ex := &Ex{
		nazwa: konf.Exchanger.Nazwa,
		kind:  konf.Exchanger.Kind,
	}

	err = ex.kanal(konf.Url)
	if err != nil {
		return nil, fmt.Errorf("błąd ex.polacz(): %v", err)
	}

	if przygotujEx, ok := rodzajeEx[konf.Exchanger.Rodzaj]; ok {
		if err := przygotujEx(ex); err != nil {
			return nil, fmt.Errorf("błąd przygotujEx(): %v", err)
		}
	} else {
		return nil, fmt.Errorf("błąd rodzajeEx[]: nieznany rodzaj exchangera")
	}

	return ex, nil
}

type Ex struct {
	ch         *amqp.Channel
	nazwa      string
	kind       string
	publikuj   func(dane []byte) error
	routingKey string
}

func (ex *Ex) kanal(url string) error {
	conn, err := kr.polaczenie(url)
	if err != nil {
		return fmt.Errorf("błąd kr.polaczenie(): %v", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("błąd conn.Channel(): %v", err)
	}
	ex.ch = ch
	return nil
}
