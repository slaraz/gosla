package v2

import (
	"encoding/json"
	"log"
	"testing"
	"time"
)

type baza struct {
	items map[int]*dbItem
}

func nowaBaza(rozmiar int) baza {
	db := baza{}
	db.items = make(map[int]*dbItem, rozmiar)
	for i := 1; i <= rozmiar; i++ {
		item := dbItem{id: i}
		db.items[item.id] = &item
	}
	return db
}

func (db baza) ids() <-chan int {
	ch := make(chan int)
	go func() {
		for k, _ := range db.items {
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
	item, ok := db.items[id]
	if !ok {
		log.Fatal("item !ok", id)
	}
	item.wyslaniaOK++
}

func (db baza) WyslanyNOK(id int) {
	item, ok := db.items[id]
	if !ok {
		log.Fatal("item !ok", id)
	}
	item.wyslaniaNOK++
}

func (db baza) Odebrany(id int) {
	item, ok := db.items[id]
	if !ok {
		log.Fatal("item !ok", id)
	}
	item.odebrania++
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
	const ile = 1000 * 1000
	db := nowaBaza(ile)
	start := time.Now()
	for id := range db.ids() {
		db.WyslanyNOK(id)
		db.Odebrany(id)
	}
	czas := time.Since(start)
	log.Printf("czas wysyłania: wszystkie = %v, jeden = %v, na sek = %.1f", czas, czas/ile, ile/czas.Seconds())
	db.DrukujStats()
}

func Test_NAPIERAJ_baze(t *testing.T) {
	const ile = 1000 * 1000
	db := nowaBaza(ile)
	// Podłączamy się do Rabbita.
	mojex := MusiExchanger(RABBIT, "moj testowy", "stdex", "fanout")
	mojqu := MusiQuełełe(RABBIT, "moja testowa", "moj testowy", "stdque", func(bajty []byte) error {
		// Tylko potwierdzenie.
		var id int
		json.Unmarshal(bajty, &id)
		db.Odebrany(id)
		return nil
	})

	// Wysyłamy do Rabbita.
	start := time.Now()
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
	mojex.Close()
	czas := time.Since(start)
	log.Printf("czas wysyłania: wszystkie = %v, jeden = %v, na sek = %.1f", czas, czas/ile, ile/czas.Seconds())
	time.Sleep(15 * time.Second)
	mojqu.Close()
	time.Sleep(time.Second)
	db.DrukujStats()
}
