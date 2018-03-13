package kostka1

import "github.com/slaraz/gosla/kostka/kolor"

type Kostka struct {
	Boki [6]*Scianka
}

type Scianka struct {
	Pola [9]kolor.Kolor
}

func NowaKostka() *Kostka {
	var kos = new(Kostka)
	for b, k := range kolor.Kolory {
		s := nowaScianka(k)
		kos.Boki[b] = s
	}
	return kos
}

func nowaScianka(k kolor.Kolor) *Scianka {
	var scia = new(Scianka)
	for i := 0; i < 9; i++ {
		scia.Pola[i] = k
	}
	return scia
}
