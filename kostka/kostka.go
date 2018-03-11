package main

type Kostka struct {
	Boki [6]*Scianka
}

type Scianka struct {
	Pola [9]Kolor
}

type Kolor byte

const (
	ziel Kolor = 1 + iota
	czer
	czar
	zolt
	nieb
	poma
)

var Kolory = []Kolor{ziel, czer, czar, zolt, nieb, poma}

func nowaKostka() *Kostka {
	var kos = new(Kostka)
	for b, k := range Kolory {
		s := nowaScianka(k)
		kos.Boki[b] = s
	}
	return kos
}

func nowaScianka(k Kolor) *Scianka {
	var scia = new(Scianka)
	for i := 0; i < 9; i++ {
		scia.Pola[i] = k
	}
	return scia
}

func (k Kostka) ObrotYGoraLewo() *Kostka {
	return nil
}
func (k Kostka) ObrotYGoraPrawo() *Kostka {
	return nil
}
