package kostka2

import (
	"github.com/slaraz/gosla/kostka/kolor"
)

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

//go:generate echo generuję ruchy
//go:generate go run gen/ruchyGen.go

// func (k Kostka) obrotYGoraPrawo2() {
// 	buf = k.Kwadraciki[0]
// 	k.Kwadraciki[0] = k.Kwadraciki[5]
// 	k.Kwadraciki[5] = k.Kwadraciki[7]
// 	k.Kwadraciki[7] = k.Kwadraciki[2]
// 	k.Kwadraciki[2] = buf

// 	buf = k.Kwadraciki[3]
// 	k.Kwadraciki[3] = k.Kwadraciki[6]
// 	k.Kwadraciki[6] = k.Kwadraciki[4]
// 	k.Kwadraciki[4] = k.Kwadraciki[1]
// 	k.Kwadraciki[1] = buf

// 	buf = k.Kwadraciki[8]
// 	k.Kwadraciki[8] = k.Kwadraciki[16]
// 	k.Kwadraciki[16] = k.Kwadraciki[24]
// 	k.Kwadraciki[24] = k.Kwadraciki[47]
// 	k.Kwadraciki[47] = buf

// 	buf = k.Kwadraciki[9]
// 	k.Kwadraciki[9] = k.Kwadraciki[17]
// 	k.Kwadraciki[17] = k.Kwadraciki[25]
// 	k.Kwadraciki[25] = k.Kwadraciki[46]
// 	k.Kwadraciki[46] = buf

// 	buf = k.Kwadraciki[10]
// 	k.Kwadraciki[10] = k.Kwadraciki[18]
// 	k.Kwadraciki[18] = k.Kwadraciki[26]
// 	k.Kwadraciki[26] = k.Kwadraciki[45]
// 	k.Kwadraciki[45] = buf
// }
