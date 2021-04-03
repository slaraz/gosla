package v2

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func MusiExchanger(url, nazwa, rodzaj, kind string) *Ex {
	ex, err := nowyEx(url, nazwa, rodzaj, kind)
	if err != nil {
		log.Fatalf("[Królik.Ex] błąd MusiExchanger(): %v", err)
	}
	return ex
}

func (ex *Ex) WyslijJSON(v interface{}) error {
	bajty, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("json.Marshal(): %v", err)
	}
	if ex.publikuj == nil {
		return fmt.Errorf("ex.pulikuj == nil")
	}
	err = ex.publikuj(bajty)
	if err != nil {
		return fmt.Errorf("ex.publikuj(): %v", err)
	}
	return nil
}

func (ex *Ex) Close() {
	if ex.sesja != nil {
		ex.sesja.Close()
	}
}

type Ex struct {
	sesja    *sesjaa
	nazwa    string
	kind     string
	publikuj func(dane []byte) error
}

func nowyEx(url, nazwa, rodzaj, kind string) (*Ex, error) {
	ex := &Ex{
		nazwa: nazwa,
		kind:  kind,
	}

	przygotujEx, ok := rozneEx[rodzaj]
	if !ok {
		return nil, fmt.Errorf("rozneEx[]: nieznany rodzaj exchangera")
	}
	przygotuj := func(chann *amqp.Channel) error {
		log.Printf("[Królik.Ex] Przygotowuję -> [%q:%s:%s]", nazwa, rodzaj, kind)
		return przygotujEx(ex, chann)
	}

	sesja := Otworz(url, przygotuj)
	ex.sesja = sesja

	return ex, nil
}
