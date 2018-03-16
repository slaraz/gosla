package main

import (
	mrand "math/rand"
	"testing"
)

func BenchmarkMathRand(b *testing.B) {
	//var i int
	for n := 0; n < b.N; n++ {
		_ = mrand.Intn(48)
	}
}
