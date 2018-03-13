package kostka2

import (
	"fmt"

	"github.com/slaraz/gosla/kostka/kolor"
)

type Kostka struct {
	Kwadraciki [54]kolor.Kolor
}

func NowaKostka() *Kostka {
	kost := new(Kostka)
	i := 0
	for _, k := range kolor.Kolory {
		for p := 0; p < 9; p++ {
			kost.Kwadraciki[i] = k
			i++
		}
	}
	return kost
}

func (k Kostka) WszystkieRuchy() []func() {
	return nil
}

func (k Kostka) String() string {
	return fmt.Sprintf("%v", k.Kwadraciki)
}
