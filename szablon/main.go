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
	a := []func(){
		func() { fmt.Println(rand.Intn(48)) },
		func() { fmt.Println(rand.Intn(48)) },
		func() { fmt.Println(rand.Intn(48)) },
		func() { fmt.Println(rand.Intn(48)) },
	}
	for _, x := range a {
		x()
		fmt.Printf("%v %t\n", x, x)
	}
}
