package main

import (
	"fmt"
	"time"

	"github.com/slaraz/gosla/kostka/kostka2"
)

func main() {
	fmt.Println("Start.")

	k := kostka2.NowaKostka()
	ruchy := k.WszystkieRuchy()

	t := time.Now()
	for i := 0; i < 1*1000*1000; i++ {
		ruchy[0]()
	}
	fmt.Println(time.Now().Sub(t))

	k.Drukuj()

}

//kostka ma po 9  jednego koloru
