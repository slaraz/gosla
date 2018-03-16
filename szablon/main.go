package main

import (
	"fmt"
	"math/rand"
)

type a struct {
	tab [2]int
}

func newA() a {
	return a{[2]int{1, 2}}
}

func (v a) mody() {
	v.tab[0] = 5
}

func main() {
	fmt.Println(rand.Intn(48))
	fmt.Println(rand.Intn(48))
	fmt.Println(rand.Intn(48))
	fmt.Println(rand.Intn(48))
}
