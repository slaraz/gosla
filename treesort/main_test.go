package main

import (
	"sort"
	"testing"
)

func BenchmarkSort(b *testing.B) {

	for n := 0; n < b.N; n++ {
		values := []int{6, 4, 6, 4, 67, 2, 76, 899, 54, 32, 4, 0}
		Sort(values)
	}
}
func BenchmarkSortSort(b *testing.B) {

	for n := 0; n < b.N; n++ {
		values := []int{6, 4, 6, 4, 67, 2, 76, 899, 54, 32, 4, 0}
		sort.Ints(values)
	}
}
