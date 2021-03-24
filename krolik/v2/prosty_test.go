package v2

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/streadway/amqp"
	"gopkg.in/yaml.v3"
)

func Test_NowyExchanger_PublikujJSON(t *testing.T) {
	MOJ_EX := `---
url: amqp://guest:guest@localhost:5672/krolik
exchanger:
  nazwa: moj testowy
  kind: fanout
  rodzaj: szybki
`

	mojex := MusiNowyExchanger(MOJ_EX)
	defer mojex.Close()

	a := struct {
		id   int
		imie string
	}{5, "ala"}

	if err := mojex.PublikujJSON(a); err != nil {
		t.Errorf("błąd PublikujJSON(): %v", err)
	}
}

// ---

var kr *krolik2

type krolik2 struct {
	conn map[string]*amqp.Connection
	exs  map[string]*Ex
}

type konfig struct {
	Url       string `yaml:"url"` // https://www.rabbitmq.com/uri-spec.html
	Exchanger struct {
		Nazwa      string `yaml:"nazwa"`
		Kind       string `yaml:"kind"`
		Rodzaj     string `yaml:"rodzaj"`
		RoutingKey string `yaml:"routingKey,omitempty"`
	} `yaml:"exchanger"`
}

func MusiNowyExchanger(conf string) *Ex {
	ex, err := nowyExchanger(conf)
	if err != nil {
		log.Fatalf("błąd NowyExchanger(): %v", err)
	}
	return ex
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

// ---

type Ex struct {
	ch         *amqp.Channel
	nazwa      string
	kind       string
	publikuj   func(dane []byte) error
	routingKey string
}

func (ex *Ex) Close() {
	ex.ch.Close()
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

// ---

var rodzajeEx = map[string]func(*Ex) error{
	"szybki": przygotujSzybki,
	//"zwykly": przygotujZwykly,
	"pewny": przygotujPewny,
}

// szybki

func przygotujSzybki(ex *Ex) error {
	err := ex.ch.ExchangeDeclare(
		ex.nazwa, // nazwa exchangera
		ex.kind,  // typ: direct, fanout, topic, headers
		true,     // durable - czy ma przerzyć reset serwera
		false,    // autodelete - czy skasować jeśli brak podłączonych kolejek
		false,    // internal
		false,    // noWait
		nil,      // arguments
	)
	if err != nil {
		return fmt.Errorf("błąd ExchangeDeclare(): %v", err)
	}
	ex.publikuj = ex.publikujSzybko
	return nil
}

func (ex *Ex) publikujSzybko(bajty []byte) error {
	err := ex.ch.Publish(
		ex.nazwa,
		ex.routingKey,
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			// tu można poszaleć! patrz inne parametry amqp.Publishing
			Body: bajty,
		})
	if err != nil {
		return fmt.Errorf("błąd ch.Publish(): %v", err)
	}
	return nil
}

// pewny

func przygotujPewny(ex *Ex) error {
	err := ex.ch.ExchangeDeclare(
		ex.nazwa, // nazwa exchangera
		ex.kind,  // typ: direct, fanout, topic, headers
		true,     // durable - czy ma przerzyć reset serwera
		false,    // autodelete - czy skasować jeśli brak podłączonych kolejek
		false,    // internal
		false,    // noWait
		nil,      // arguments
	)
	if err != nil {
		return fmt.Errorf("błąd ExchangeDeclare(): %v", err)
	}
	return nil
}
