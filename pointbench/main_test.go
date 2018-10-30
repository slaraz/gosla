package main

import "testing"

type Adder struct {
	x, y int
}

func (a *Adder) add() int {
	return a.x + a.y
}
func BenchmarkWithoutPointer(b *testing.B) {
	accum := 0
	for i := 0; i < b.N; i++ {
		adder := Adder{accum, i}
		accum = adder.add()
	}
	_ = accum
}

func BenchmarkWithPointer(b *testing.B) {
	accum := 0
	for i := 0; i < b.N; i++ {
		adder := &Adder{accum, i}
		accum = adder.add()
	}
	_ = accum
}
