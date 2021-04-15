package v2

import (
	"encoding/json"
	"log"
	"testing"
	"time"
)

/*

# serwer:
docker run -d -p 15672:15672 -p 5672:5672 rabbitmq:3-management
# vhost krolik:
curl -u guest:guest -X PUT http://localhost:15672/api/vhosts/krolik
# skasuj vhost krolik:
curl -u guest:guest -X DELETE http://localhost:15672/api/vhosts/krolik

*/
var RABBIT = "amqp:///krolik"
var EX = "*** moj testowy ***"
var QUE = "*** moja testowa ***"

func Test_Wyslij_I_Odbierz(t *testing.T) {
	// Przygotowujemy Exchanger i Quełełe.
	mojex := MusiExchanger(RABBIT, EX, "stdex", "fanout")
	defer mojex.Close()
	mojqu := MusiQuełełe(RABBIT, queParam{QUE, EX}, "stdque")
	defer mojqu.Close()

	// Odbieramy wiadomość.
	mojqu.Konsumuj(odbieranie)

	// Wysyłamy wiadomość
	wiad := wiadomoscTestowa()
	if err := mojex.WyslijJSON(wiad); err != nil {
		t.Fatalf("błąd WyslijJSON: %v", err)
	}
	log.Printf("* wysłałem %q", wiad.Id)

	// Czekamy na odebranie wszystkiego z kolejki.
	<-mojqu.Pusta

}

func Test_NABIJ_1000(t *testing.T) {
	mojex := MusiExchanger(RABBIT, EX, "stdex", "fanout")
	defer mojex.Close()

	start := time.Now()
	for i := 0; i < 1000; i++ {
		if err := mojex.WyslijJSON(wiadomoscNumerowana()); err != nil {
			log.Printf("błąd WyslijJSON: %v", err)
		}
	}
	log.Printf("nabiłem 1000 w %v", time.Since(start))
}

func Test_Quełełe_KOŃCZENIE(t *testing.T) {
	// Połączenie się z kolejką w RabbitMQ.
	mojqu := MusiQuełełe(RABBIT, queParam{QUE, EX}, "stdque")
	mojqu.Konsumuj(odbieranie)

	// Czas na odebranie.
	time.Sleep(10 * time.Millisecond)

	// Zamykanie połączenia.
	log.Println("----KOŃCZENIE-------------")
	start := time.Now()
	mojqu.Close()
	log.Println("----KONIEC----", time.Since(start))
}

func Test_CIURKAJ_po_malutku(t *testing.T) {
	// Podłączamy się do Rabbita.
	mojex := MusiExchanger(RABBIT, EX, "stdex", "fanout")
	defer mojex.Close()
	mojqu := MusiQuełełe(RABBIT, queParam{QUE, EX}, "stdque")
	mojqu.Konsumuj(odbieranie)
	defer mojqu.Close()

	// Wysyłanie do Rabbita.
	for {
		time.Sleep(time.Second * 1)
		wiad := wiadomoscTestowa()
		err := mojex.WyslijJSON(wiad)
		if err != nil {
			continue
		}
		log.Printf("* wysłałem %q", wiad.Id)
	}
}

func Test_NAPIERAJ_ile_wlezie(t *testing.T) {
	ileWlezie := 100 * 1000
	// Podłączamy się do Rabbita.
	mojex := MusiExchanger(RABBIT, EX, "stdex", "fanout")
	mojqu := MusiQuełełe(RABBIT, queParam{QUE, EX}, "stdque")

	ileOdebr := 0
	mojqu.Konsumuj(func(bajty []byte) error {
		ileOdebr++
		return nil
	})

	// Wysyłamy do Rabbita.
	ileWys := 0
	for i := 0; i < ileWlezie; i++ {
		if err := mojex.WyslijJSON(wiadomoscTestowa()); err != nil {
			log.Printf("błąd WyslijJSON: %v", err)
			time.Sleep(time.Second)
		}
		ileWys++
		if ileWys%1e4 == 0 {
			log.Printf("wysłałem %dk", ileWys/1e3)
		}
	}
	time.Sleep(time.Second)
	mojex.Close()
	time.Sleep(time.Second)
	mojqu.Close()
	log.Printf("wysłałem %d, odebrałem %d", ileWys, ileOdebr)
}

// ---

func odbieranie(body []byte) error {
	var wiad wiadomosc
	err := json.Unmarshal(body, &wiad)
	if err != nil {
		log.Fatalf("błąd json.Unmarshal(): %v", err)
	}
	log.Printf("* odebrałem %v", wiad)
	return nil
}
