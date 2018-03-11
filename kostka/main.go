package main

import "fmt"

func main() {
	fmt.Println("Start.")
}

type kostka struct {
	boki [6]*scianka
}

type scianka struct {
	pola [9]kolor
}

type kolor int

const (
	ziel kolor = iota
	czer
	czar
	zolt
	nieb
	poma
)
