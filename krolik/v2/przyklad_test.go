package v2

import (
	"log"
	"testing"
	"time"
)

// # serwer:
// docker run -d -p 15672:15672 -p 5672:5672 rabbitmq:3-management
// # utworz vhost krolik:
// curl -u guest:guest -X PUT http://localhost:15672/api/vhosts/krolik
// # skasuj vhost krolik:
// curl -u guest:guest -X DELETE http://localhost:15672/api/vhosts/krolik
var RABBIT = "amqp://guest:guest@localhost:5672/krolik"

func Test_NowyExchanger_PublikujJSON(t *testing.T) {
	// Podłączamy się do Rabbita.
	mojex := MusiExchanger(RABBIT, "moj testowy", "stdex", "fanout")
	defer mojex.Close()

	mojqu := MusiQuełełe(RABBIT, "moja testowa", "moj testowy", "stdque", odbieranie)
	defer mojqu.Close()

	// Wysyłamy do Rabbita.
	type dane struct {
		Id   int
		Imie string
	}
	wys := dane{9, "Ala"}
	if err := mojex.WyslijJSON(wys); err != nil {
		t.Fatalf("błąd WyslijJSON: %v", err)
	}

	// Czas na odebranie.
	time.Sleep(50 * time.Millisecond)
}

var odbieranie = func(bajty []byte) error {
	log.Println("* odbieranie", string(bajty), time.Since(start))
	time.Sleep(time.Millisecond)
	return nil
}

func Test_KOŃCZENIE(t *testing.T) {

	mojqu := MusiQuełełe(RABBIT, "moja testowa", "moj testowy", "stdque", odbieranie)

	// Czas na odebranie.
	time.Sleep(5 * time.Millisecond)
	log.Println("----KOŃCZENIE-------------")
	start = time.Now()
	mojqu.Close()
	log.Println("----KONIEC----", time.Since(start))
	time.Sleep(time.Second)
}

func Test_NabijanieRabbita(t *testing.T) {
	// Podłączamy się do Rabbita.
	mojex := MusiExchanger(RABBIT, "moj testowy", "stdex", "fanout")
	defer mojex.Close()

	odbieranie := func(bajty []byte) error {
		time.Sleep(time.Millisecond)
		return nil
	}
	mojqu := MusiQuełełe(RABBIT, "moja testowa", "moj testowy", "stdque", odbieranie)
	defer mojqu.Close()

	// Wysyłamy do Rabbita.
	type dane struct {
		Id   int
		Imie string
	}
	wys := dane{8, "Ula"}

	j := 0
	for {
		//time.Sleep(time.Second * 1)
		if err := mojex.WyslijJSON(wys); err != nil {
			log.Printf("błąd WyslijJSON: %v", err)
			time.Sleep(time.Second)
		}
		j++
		if j%1e4 == 0 {
			log.Printf("wysłałem %d", j/1e4)
		}
	}
}
