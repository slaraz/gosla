package v2

import (
	"fmt"
	"log"
	"testing"

	"github.com/streadway/amqp"
	"gopkg.in/yaml.v3"
)

func Test_NowyExchanger_PublikujJSON(t *testing.T) {
	MOJ_EX := `---
url: amqp://guest:guest@localhost:5672/test_krolika2
exchanger:
  nazwa: moj testowy
  kind: funout
  rodzaj: szbyki
`

	mojex := MusiNowyExchanger(MOJ_EX)
	defer mojex.Close()

	a := struct {
		id   int
		imie string
	}{5, "ala"}

	mojex.PublikujJSON(a)
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
		Nazwa  string `yaml:"nazwa"`
		Kind   string `yaml:"kind"`
		Rodzaj string `yaml:"rodzaj"`
	} `yaml:"exchanger"`
}

type Ex struct {
	ch       *amqp.Channel
	nazwa    string
	kind     string
	publikuj func(dane []byte) error
}

func MusiNowyExchanger(conf string) *Ex {
	ex, err := NowyExchanger(conf)
	if err != nil {
		log.Fatalf("błąd NowyExchanger(): %v", err)
	}
	return ex
}

func NowyExchanger(conf string) (*Ex, error) {
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
		kind:  konf.Exchanger.Nazwa,
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

	return nil, nil
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

func (ex *Ex) PublikujJSON(v interface{}) {

}

// ---

var rodzajeEx = map[string]func(*Ex) error{
	"szybki": przygotujSzybki,
	"pewny":  przygotujPewny,
}

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
	ex.publish = ex.publikujSzybko
	return nil
}

func (ex *Ex) publikujSzybko(dane []byte) error {
	err := ex.ch.Publish(
		"logs", // exchange
		"",     // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
}

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
