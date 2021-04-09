package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/slaraz/gosla/krolik/v2"
)

var RABBIT = "amqp://guest:guest@localhost:5672/krolik"

func main() {
	mojex := v2.MusiExchanger(RABBIT, "moj testowy", "stdex", "fanout")
	mojqu := v2.MusiQuełełe(RABBIT, "moja testowa", "moj testowy", "stdque", odbieranie)

	// Wysyłamy do Rabbita.
	wys := struct {
		UID  string
		Czas time.Time
	}{RandomString(3), time.Now()}
	if err := mojex.WyslijJSON(wys); err != nil {
		log.Fatalf("błąd WyslijJSON: %v", err)
	}
	log.Printf("* wysłałem %q", wys.UID)

	// Czas na odebranie.
	time.Sleep(50 * time.Millisecond)
	mojqu.Close()
	mojex.Close()
}

var odbieranie = func(bajty []byte) error {
	d := struct {
		UID  string
		Czas time.Time
	}{}
	err := json.Unmarshal(bajty, &d)
	if err != nil {
		log.Fatalf("błąd json.Unmarshal(): %v", err)
	}
	log.Printf("* odebrałem %q: %v", d.UID, time.Since(d.Czas))
	time.Sleep(time.Millisecond)
	return nil
}

func RandomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
