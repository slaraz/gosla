package main

import (
	"testing"

	"github.com/slaraz/gosla/kostka/kostka2"
)

func BenchmarkLosowyRuch(b *testing.B) {
	k := kostka2.NowaKostka()
	r := k.WszystkieRuchy()
	for n := 0; n < b.N; n++ {
		losowyRuch(r)
	}
}

func BenchmarkLosowyRuchGen(b *testing.B) {
	k := kostka2.NowaKostka()
	losRuch := genFuncLosowyRuch(k)
	for n := 0; n < b.N; n++ {
		losRuch()
	}
}
