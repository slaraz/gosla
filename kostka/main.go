package main

import (
	"fmt"

	"github.com/slaraz/gosla/kostka/kostka2"
)

func main() {
	fmt.Println("Start.")

	k := kostka2.NowaKostka()
	ruchy := k.WszystkieRuchy()

	fmt.Println("Ile ruchów:", len(ruchy))

	fmt.Println(k)
}
