package kolor

type Kolor byte

const (
	czar Kolor = 1 + iota
	ziel
	czer
	nieb
	zolt
	poma
)

//var Kolory = []Kolor{czar, czer, nieb, zolt, żółt, poma}

var nazwy = map[Kolor]string{
	czar: "czar",
	czer: "czer",
	nieb: "nieb",
	zolt: "zolt",
	poma: "poma",
	ziel: "ziel",
}

func (k Kolor) String() string {
	return nazwy[k]
}
