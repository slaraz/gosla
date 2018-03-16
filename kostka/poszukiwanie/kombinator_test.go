package poszukiwanie

import (
	"testing"

	"github.com/slaraz/gosla/kostka/kostka2"
)

func Test_kombinuj(t *testing.T) {
	k := kostka2.NowaKostka()
	Mieszaj(k.WszystkieRuchy(), 3)
	kom := &kombinator{maxpoziom: 2}
	kom.kombinuj(k, 0)
	m := kom.skrajeWyniki
	if len(m) != 144 {
		t.Fail()
	}
}

func BenchmarkKombinuj(b *testing.B) {
	for n := 0; n < b.N; n++ {
		k := kostka2.NowaKostka()
		Mieszaj(k.WszystkieRuchy(), 30)
		kom := &kombinator{maxpoziom: 3}
		kom.kombinuj(k, 0)
	}
}
