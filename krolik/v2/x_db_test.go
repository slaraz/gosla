package v2

import (
	"encoding/json"
	"log"
	"testing"
	"time"
)

type baza struct {
	items []*dbItem
}

func nowaBaza(rozmiar int) baza {
	db := baza{}
	db.items = make([]*dbItem, rozmiar)
	for i := 0; i < rozmiar; i++ {
		item := dbItem{id: i}
		db.items[i] = &item
	}
	return db
}

func (db baza) ids() <-chan int {
	ch := make(chan int)
	go func() {
		for k := range db.items {
			ch <- k
		}
		close(ch)
	}()
	return ch
}

type dbItem struct {
	id          int
	wyslaniaOK  int
	wyslaniaNOK int
	odebrania   int
}

func (item dbItem) poprawny() bool {
	return item.wyslaniaOK == 1 && item.wyslaniaNOK == 0 && item.odebrania == 1
}

func (db baza) WyslanyOK(id int) {
	db.items[id].wyslaniaOK++
}

func (db baza) WyslanyNOK(id int) {
	db.items[id].wyslaniaNOK++
}

func (db baza) Odebrany(id int) {
	db.items[id].odebrania++
}

func (db baza) DrukujStats() {
	log.Print("db stats:")
	var popr, niePopr int
	var nieOdebrane, odebraneRaz, odebraneWiele int
	var nieWyslaneOK, wyslaneOKRaz, wyslaneOKWiele int
	var nieWyslaneNOK, wyslaneNOKRaz, wyslaneNOKWiele int

	for _, item := range db.items {

		if item.odebrania == 0 {
			nieOdebrane++
		} else if item.odebrania == 1 {
			odebraneRaz++
		} else {
			odebraneWiele++
		}

		if item.wyslaniaOK == 0 {
			nieWyslaneOK++
		} else if item.wyslaniaOK == 1 {
			wyslaneOKRaz++
		} else {
			wyslaneOKWiele++
		}

		if item.wyslaniaNOK == 0 {
			nieWyslaneNOK++
		} else if item.wyslaniaNOK == 1 {
			wyslaneNOKRaz++
		} else {
			wyslaneNOKWiele++
		}

		if item.poprawny() {
			popr++
			continue
		}
		niePopr++
		//log.Print(item)
	}
	druk("ideal", popr)
	druk("nieIdeal", niePopr)
	druk("nieOdebrane", nieOdebrane)
	druk("odebraneRaz", odebraneRaz)
	druk("odebraneWiele", odebraneWiele)
	druk("nieWyslaneOK", nieWyslaneOK)
	druk("wyslaneOKRaz", wyslaneOKRaz)
	druk("wyslaneOKWiele", wyslaneOKWiele)
	druk("wyslaneNOKRaz", wyslaneNOKRaz)
	druk("wyslaneNOKWiele", wyslaneNOKWiele)
}

func druk(txt string, ile int) {
	if ile != 0 {
		log.Printf("  %s: %d", txt, ile)
	}
}

func Test_GOŁA_db(t *testing.T) {
	const ile = 200 * 1000
	db := nowaBaza(ile)
	start := time.Now()
	for id := range db.ids() {
		db.WyslanyOK(id)
		db.Odebrany(id)
	}
	printSzybkosc(start, ile)
	db.DrukujStats()
}

func Test_NAPIERAJ_szybkie(t *testing.T) {
	const ile = 1000 * 1000
	db := nowaBaza(ile)
	// Podłączamy się do Rabbita.

	mojex := MusiExchanger(RABBIT, EX, "szybki", "fanout")
	mojqu := MusiQuełełe(RABBIT, queParam{QUE, EX}, "szybka")
	mojqu.UsunWiadomosci()
	mojqu.Konsumuj(func(bajty []byte) error {
		// Tylko potwierdzenie.
		var id int
		json.Unmarshal(bajty, &id)
		db.Odebrany(id)
		return nil
	})

	// Wysyłamy do Rabbita.
	start := time.Now()
	wyslij(db, mojex)

	mojex.Close()
	printSzybkosc(start, ile)
	// Czekamy na odebranie wszystkiego z kolejki.
	<-mojqu.Pusta
	mojqu.Close()
	db.DrukujStats()
}

func Test_NAPIERAJ_std(t *testing.T) {
	const ile = 500 * 1000
	db := nowaBaza(ile)
	// Podłączamy się do Rabbita.
	mojex := MusiExchanger(RABBIT, EX, "stdex", "fanout")
	mojqu := MusiQuełełe(RABBIT, queParam{QUE, EX}, "stdque")
	mojqu.UsunWiadomosci()
	mojqu.Konsumuj(func(bajty []byte) error {
		var id int
		json.Unmarshal(bajty, &id)
		db.Odebrany(id)
		return nil
	})

	// Wysyłamy do Rabbita.
	start := time.Now()
	wyslij(db, mojex)

	mojex.Close()
	printSzybkosc(start, ile)
	// Czekamy na odebranie wszystkiego z kolejki.
	<-mojqu.Pusta
	mojqu.Close()
	db.DrukujStats()
}

// ---

func wyslij(db baza, mojex *Ex) {
	x := 0
	for id := range db.ids() {
		if err := mojex.WyslijJSON(id); err != nil {
			db.WyslanyNOK(id)
			log.Printf("błąd WyslijJSON: %v", err)
			time.Sleep(time.Second)
		} else {
			db.WyslanyOK(id)
		}
		x++
		if x%1e5 == 0 {
			log.Printf("wysłałem %dk", x/1e3)
		}
	}
}

func printSzybkosc(start time.Time, ile int64) {
	czas := time.Since(start)
	ileK := ile / 1000
	log.Printf("Szybkość: czas wysyłania: %dk-> %v, jeden-> %v, %.1fk/s", ileK, czas, time.Duration(int64(czas)/ile), float64(ileK)/czas.Seconds())
}
