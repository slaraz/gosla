package main

import (
	"fmt"
)

type a struct {
	tab []int
}

func newA() a {
	return a{[]int{1, 2}}
}

func (v a) mody() {
	v.tab[0] = 5
}

func main() {
	v := newA()
	fmt.Printf("%v\n", v)
	v.mody()
	fmt.Printf("%v\n", v)
}
