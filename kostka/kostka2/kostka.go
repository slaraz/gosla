package kostka2

import (
	"github.com/slaraz/gosla/kostka/kolor"
)

type Kostka struct {
	Kwadraciki [48]kolor.Kolor
}

func NowaKostka() *Kostka {
	kost := &Kostka{}
	i := 0
	for k := 1; k <= 6; k++ {
		for p := 0; p < 8; p++ {
			kost.Kwadraciki[i] = kolor.Kolor(k)
			i++
		}
	}
	return kost
}

func (kost *Kostka) IleNaMiejscu() (ile int) {
	for k := 1; k <= 6; k++ {
		for p := 0; p < 8; p++ {
			if kost.Kwadraciki[i] == kolor.Kolor(k) {
				ile++
			}
		}
	}
	return ile
}
