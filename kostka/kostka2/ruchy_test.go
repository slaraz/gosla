package kostka2

import (
	"math/rand"
	"testing"
)

func TestKostka_4Obroty(t *testing.T) {
	k := NowaKostka()
	c := *k
	r := k.WszystkieRuchy()
	for _, obr := range r {
		obr()
		if *k == c {
			t.Fail()
		}
		obr()
		obr()
		obr()
		if *k != c {
			t.Fail()
		}
	}
}
func TestKostka_ObrotyAB(t *testing.T) {
	k := NowaKostka()
	c := *k
	r := k.WszystkieRuchy()
	for i := 0; i < len(r); i += 2 {
		obrA := r[i]
		obrB := r[i+1]
		obrA()
		if *k == c {
			t.Fail()
		}
		obrB()
		if *k != c {
			t.Fail()
		}
	}
}

func BenchmarkObrotKlasyk(b *testing.B) {
	k := NowaKostka()
	for n := 0; n < b.N; n++ {
		k.obrotYGoraPrawo()
	}
}
func BenchmarkObrotNaPiechote(b *testing.B) {
	k := NowaKostka()
	for n := 0; n < b.N; n++ {
		k.ObrZTylB()
	}
}
func BenchmarkNowaKostka(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = NowaKostka()
	}
}
func BenchmarkPorównanie(b *testing.B) {
	k1 := NowaKostka()
	k2 := NowaKostka()
	for n := 0; n < b.N; n++ {
		_ = *k1 == *k2
	}
}

func BenchmarkMathRand(b *testing.B) {
	//var i int
	for n := 0; n < b.N; n++ {
		_ = rand.Intn(48)
	}
}
