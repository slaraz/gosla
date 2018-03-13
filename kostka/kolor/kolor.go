package kolor

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

func (k Kolor) String() string {
	var w string
	switch k {
	case ziel:
		w = "ziel"
	case czer:
		w = "czer"
	case czar:
		w = "czar"
	case zolt:
		w = "zolt"
	case nieb:
		w = "nieb"
	case poma:
		w = "poma"
	}
	return w
}
