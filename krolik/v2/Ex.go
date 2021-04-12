package v2

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type Ex struct {
	sesja    *sesjaa
	nazwa    string
	kind     string
	publikuj func(dane []byte) error
}

func MusiExchanger(url, nazwa, rodzaj, kind string) *Ex {
	ex, err := nowyEx(url, nazwa, rodzaj, kind)
	if err != nil {
		log.Fatalf("błąd MusiExchanger(): %v", err)
	}
	return ex
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
	przygotuj := func(chann *amqp.Channel, log *log.Logger) error {
		log.Printf("Przygotowuję [%s->%s]", rodzaj, kind)
		return przygotujEx(ex, chann)
	}

	nazwaSesji := fmt.Sprintf("EX(%s)", nazwa)

	sesja := otworz(url, przygotuj, nazwaSesji)
	ex.sesja = sesja

	return ex, nil
}

func (ex *Ex) WyslijJSON(v interface{}) error {
	if ex.publikuj == nil {
		return fmt.Errorf("ex.publikuj == nil")
	}
	if !ex.sesja.czyOK {
		return fmt.Errorf("brak połączenia")
	}
	if ex.sesja.czyBlock {
		return fmt.Errorf("połączenie blokowane przez serwer")
	}
	if !ex.sesja.czyFlow {
		return fmt.Errorf("serwer prosi o łaskę (Flow)")
	}

	bajty, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("json.Marshal(): %v", err)
	}
	err = ex.publikuj(bajty)
	if err != nil {
		return fmt.Errorf("ex.publikuj(): %v", err)
	}
	return nil
}

func (ex *Ex) Close() {
	if ex.sesja != nil {
		ex.sesja.close()
	}
}
