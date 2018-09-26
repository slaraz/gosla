package poszukiwanie

import "github.com/slaraz/gosla/kostka/kostka2"

func kombinuj(k *kostka2.Kostka) []kostka2.Kostka {
	kom := &kombinator{maxpoziom: 3}
	kom.kombinuj(k, 0)
	return kom.skrajeWyniki
}

type kombinator struct {
	maxpoziom    int
	skrajeWyniki []kostka2.Kostka
}

func (kk *kombinator) kombinuj(k *kostka2.Kostka, poziom int) {
	poziom++
	for _, r := range k.WszystkieRuchy() {
		r()
		czyUlozona(k)
		if poziom < kk.maxpoziom {
			kk.kombinuj(k, poziom)
		} else {
			kk.skrajeWyniki = append(kk.skrajeWyniki, *k)
		}
	}
}
