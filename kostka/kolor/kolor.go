package kolor

type Kolor byte

const (
	Czar Kolor = 1 + iota
	Ziel
	Czer
	Nieb
	Zolt
	Poma
)

//var Kolory = []Kolor{czar, czer, nieb, zolt, żółt, poma}

var nazwy = map[Kolor]string{
	Czar: "Czar",
	Ziel: "Ziel",
	Czer: "Czer",
	Nieb: "Nieb",
	Zolt: "Zolt",
	Poma: "Poma",
}

func (k Kolor) String() string {
	return nazwy[k]
}
