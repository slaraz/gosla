package v2

import (
	"log"
	"testing"
	"time"
)

var RABBIT = "amqp://guest:guest@localhost:5672/krolik"

func Test_NowyExchanger_PublikujJSON(t *testing.T) {
	// Podłączamy się do Rabbita.
	mojex := MusiExchanger(RABBIT, "moj testowy", "std", "fanout")
	defer mojex.Close()

	mojqu := MusiQuełełe(RABBIT, "moja testowa", "xxx")
	defer mojqu.Close()

	// Wysyłamy do Rabbita.
	type dane struct {
		Id   int
		Imie string
	}
	wys := dane{5, "Ala"}
	if err := mojex.WyslijJSON(wys); err != nil {
		t.Fatalf("błąd WyslijJSON: %v", err)
	}

	// Odbieramy z Rabbita.
	odb := dane{}
	if err := mojqu.OdbierzJSON(&odb); err != nil {
		t.Fatalf("błąd OdbierzJSON: %v", err)
	}

	// Sprawdzamy wynik.
	if wys != odb {
		t.Fatalf("błąd królika:\nwysłałem: %v\nodebrałem: %v", wys, odb)
	}
}

func Test_WyslijJSON_NieskonczonaPetla(t *testing.T) {
	// Podłączamy się do Rabbita.
	mojex := MusiExchanger(RABBIT, "moj testowy", "std", "fanout")
	defer mojex.Close()

	// Wysyłamy do Rabbita.
	type dane struct {
		Id   int
		Imie string
	}
	wys := dane{6, "Ula"}

	for {
		time.Sleep(time.Second * 3)
		if err := mojex.WyslijJSON(wys); err != nil {
			log.Printf("błąd WyslijJSON: %v", err)
		} else {
			log.Println("Wysyłanie ok.")
		}
	}
}
