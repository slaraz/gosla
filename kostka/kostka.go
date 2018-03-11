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

func (k Kostka) ObrotYGoraLewo() {
	return
}
func (k Kostka) ObrotYGoraPrawo() {
	return
}

// ObrotYGoraLewo() *IKostka
// ObrotYGoraPrawo() *IKostka
// ObrotYSrodekLewo() *IKostka
// ObrotYSrodekPrawo() *IKostka
// ObrotYDolLewo() *IKostka
// ObrotYDolPrawo() *IKostka

// ObrotXLewoGora() *IKostka
// ObrotXLewoDol() *IKostka
// ObrotXSrodekGora() *IKostka
// ObrotXSrodekDol() *IKostka
// ObrotXPrawoGora() *IKostka
// ObrotXPrawoDol() *IKostka

// ObrotZPrzodPrawo() *IKostka
// ObrotZPrzodLewo() *IKostka
// ObrotZSrodekPrawo() *IKostka
// ObrotZSrodekLewo() *IKostka
// ObrotZTylPrawo() *IKostka
// ObrotZTylLewo() *IKostka
