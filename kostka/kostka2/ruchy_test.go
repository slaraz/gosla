package kostka2

import (
	"fmt"
	"testing"

	"github.com/slaraz/gosla/kostka/kolor"
)

func TestKostka_ObrótXGóraDół(t *testing.T) {
	tests := []struct {
		name string
		k    *Kostka
		i    int
		w    kolor.Kolor
	}{
		{"jeden obrót",
			NowaKostka(),
			17,
			kolor.Nieb},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.k.ObrótXGóraDół()
			fmt.Println(tt.k.Kwadraciki[tt.i])
			if tt.k.Kwadraciki[tt.i] != tt.w {
				t.Fail()
			}
		})
	}
}

func BenchmarkObrótXGóraDół(b *testing.B) {
	k := NowaKostka()
	for n := 0; n < b.N; n++ {
		k.ObrótXGóraDół()
	}
}
func BenchmarkObrotGD(b *testing.B) {
	k := NowaKostka()
	for n := 0; n < b.N; n++ {
		k.obrotYGoraPrawo()
	}
}
