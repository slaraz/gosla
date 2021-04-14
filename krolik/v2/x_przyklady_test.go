package v2

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
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

func Test_Wyslij_I_Odbierz(t *testing.T) {
	// Podłączamy się do Rabbita.
	mojex := MusiExchanger(RABBIT, "moj testowy", "stdex", "fanout")
	defer mojex.Close()

	// Wysyłamy wiadomość do Rabbita.
	wiad := wiadomoscTestowa()
	if err := mojex.WyslijJSON(wiad); err != nil {
		t.Fatalf("błąd WyslijJSON: %v", err)
	}
	log.Printf("* wysłałem %q", wiad.Id)

	// Podłączenie do odbierania.
	mojqu := MusiQuełełe(RABBIT, queParam{"moja testowa", "moj testowy"}, "stdque")
	defer mojqu.Close()
	mojqu.Konsumuj(odbieranie)

	// Czas na odebranie wiadomości przez "go_fera odbieranie()"
	time.Sleep(100 * time.Millisecond)
	// Szukaj "* odebrałem" na konsoli.
}

func Test_NABIJ_1000(t *testing.T) {
	mojex := MusiExchanger(RABBIT, "moj testowy", "stdex", "fanout")
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
	mojqu := MusiQuełełe(RABBIT, queParam{"moja testowa", "moj testowy"}, "stdque")
	mojqu.Konsumuj(odbieranie)

	// Czas na odebranie.
	time.Sleep(10 * time.Millisecond)

	// Zamykanie połączenia.
	log.Println("----KOŃCZENIE-------------")
	start := time.Now()
	mojqu.Close()
	log.Println("----KONIEC----", time.Since(start))
}

// func Test_CIURKAJ_po_malutku(t *testing.T) {
// 	// Podłączamy się do Rabbita.
// 	mojex := MusiExchanger(RABBIT, "moj testowy", "stdex", "fanout")
// 	defer mojex.Close()
// 	mojqu := MusiQuełełe(RABBIT, "moja testowa", "moj testowy", "stdque", odbieranie)
// 	defer mojqu.Close()

// 	// Wysyłanie do Rabbita.
// 	for {
// 		time.Sleep(time.Second * 1)
// 		wiad := wiadomoscTestowa()
// 		err := mojex.WyslijJSON(wiad)
// 		if err != nil {
// 			continue
// 		}
// 		log.Printf("* wysłałem %q", wiad.Id)
// 	}
// }

// func Test_NAPIERAJ_ile_wlezie(t *testing.T) {
// 	ileWlezie := 100 * 1000
// 	// Podłączamy się do Rabbita.
// 	mojex := MusiExchanger(RABBIT, "moj testowy", "stdex", "fanout")
// 	ileOdebr := 0
// 	mojqu := MusiQuełełe(RABBIT, "moja testowa", "moj testowy", "stdque", func(bajty []byte) error {
// 		// Tylko potwierdzenie.
// 		ileOdebr++
// 		return nil
// 	})

// 	// Wysyłamy do Rabbita.
// 	ileWys := 0
// 	for i := 0; i < ileWlezie; i++ {
// 		//time.Sleep(10*time.Microsecond)
// 		wiad := wiadomoscTestowa()
// 		if err := mojex.WyslijJSON(wiad); err != nil {
// 			log.Printf("błąd WyslijJSON: %v", err)
// 			time.Sleep(time.Second)
// 		}
// 		ileWys++
// 		if ileWys%1e4 == 0 {
// 			log.Printf("wysłałem %dk", ileWys/1e3)
// 		}
// 	}
// 	time.Sleep(time.Second)
// 	mojex.Close()
// 	time.Sleep(time.Second)
// 	mojqu.Close()
// 	log.Printf("wysłałem %d, odebrałem %d", ileWys, ileOdebr)
// }

// func Test_ODBIERAJ_po_malutku(t *testing.T) {
// 	// Podłączamy się do Rabbita.
// 	mojqu := MusiQuełełe(RABBIT, "moja testowa", "moj testowy", "stdque", func(body []byte) error {
// 		time.Sleep(1000 * time.Millisecond)
// 		return odbieranie(body)
// 	})
// 	defer mojqu.Close()

// 	select {}
// }

// func Test_ODBIERZ_jeden(t *testing.T) {
// 	// Podłączamy się do Rabbita.
// 	odb := make(chan []byte)
// 	mojqu := MusiQuełełe(RABBIT, "moja testowa", "moj testowy", "stdque", func(body []byte) error {
// 		odb <- body
// 		return nil
// 	})
// 	odbieranie(<-odb)
// 	mojqu.Close()
// }

// ---

func odbieranie(body []byte) error {
	wiad := wiadomoscPusta()
	err := json.Unmarshal(body, &wiad)
	if err != nil {
		log.Fatalf("błąd json.Unmarshal(): %v", err)
	}
	log.Printf("* odebrałem %v", wiad)
	return nil
}

// --- wiadomosc

type wiadomosc struct {
	Id   string
	Czas time.Time
}

func wiadomoscTestowa() wiadomosc {
	return wiadomosc{
		Id:   RandomString(3),
		Czas: time.Now(),
	}
}

var kolejny int = 0
var sesja string

func wiadomoscNumerowana() wiadomosc {
	if sesja == "" {
		sesja = RandomString(3)
	}
	kolejny++
	return wiadomosc{
		Id:   fmt.Sprintf("%s:%04d", sesja, kolejny),
		Czas: time.Now(),
	}
}

func wiadomoscPusta() (pusta wiadomosc) { return }

func (wiad wiadomosc) String() string {
	if wiad == wiadomoscPusta() {
		return "PUSTA WIAD"
	} else {
		return fmt.Sprintf("%q: %v", wiad.Id, time.Since(wiad.Czas))
	}
}

// --- random string

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
