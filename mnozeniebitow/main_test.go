package main

import (
	"fmt"
	"testing"
)

const n = 2000

type iii int64

func BenchmarkMnozenie64(t *testing.B) {
	a := make([]iii, n)
	b := make([]iii, n)
	fmt.Println("x")
	for i := 0; i < t.N; i++ {
		// mnozenie(a, b)
		for j := 0; j < n; j++ {
			c := a[j] | b[j]
			_ = c
		}
	}

}

func mnozenie(a, b []iii) {

}
