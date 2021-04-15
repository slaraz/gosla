package v2

import (
	"fmt"
	"math/rand"
	"time"
)

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

var numerKolejnyWiadomosci int = 0
var unikalnaNazwaSesji string

func wiadomoscNumerowana() wiadomosc {
	if unikalnaNazwaSesji == "" {
		unikalnaNazwaSesji = RandomString(3)
	}
	numerKolejnyWiadomosci++
	return wiadomosc{
		Id:   fmt.Sprintf("%s:%04d", unikalnaNazwaSesji, numerKolejnyWiadomosci),
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
