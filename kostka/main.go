package main

import (
	"fmt"

	"github.com/slaraz/gosla/kostka/kostka2"
	"github.com/slaraz/gosla/kostka/poszukiwanie"
)

func main() {
	fmt.Println("Start...")

	k := kostka2.NowaKostka()
	r := k.WszystkieRuchy()
	poszukiwanie.Mieszaj(r, 4)
	k.Drukuj()

	poszukiwanie.Szukaj(k)
}
