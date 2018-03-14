package kostka2

import (
	"github.com/slaraz/gosla/kostka/kolor"
)

func (k *Kostka) WszystkieRuchy() []func() {
	return []func(){
		k.obrotYGoraPrawo,
	}
}

var r = [][]int{
	{0, 5, 7, 2},
	{3, 6, 4, 1},
	{8, 16, 24, 47},
	{9, 17, 25, 46},
	{10, 18, 26, 45},
}
var buf kolor.Kolor
var i, j int

func (k *Kostka) obrotYGoraPrawo() {
	for i = 0; i < 5; i++ {
		buf = k.Kwadraciki[r[i][3]]
		for j = 3; j > 0; j-- {
			k.Kwadraciki[r[i][j]] = k.Kwadraciki[r[i][j-1]]
		}
		k.Kwadraciki[r[i][0]] = buf
	}
}
