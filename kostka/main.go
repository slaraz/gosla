package main

import (
	"fmt"
	"time"

	"github.com/slaraz/gosla/kostka/kostka2"
)

func main() {
	fmt.Println("Start.")

	k := kostka2.NowaKostka()
	ruch := k.WszystkieRuchy()[1]

	ta := time.After(1 * time.Second)

	i := 0
loop:
	for {

		ruch()
		i++

		if i%1000000 == 0 {
			select {
			case <-ta:
				break loop
			default:
			}
		}
	}

	fmt.Println("1 sek.:", i, "ruchów")

	t := time.Now()
	for i = 0; i < 100*1000*1000; i++ {
		ruch()
	}
	fmt.Println(time.Now().Sub(t))

	//	k.Drukuj()

}

//kostka ma po 9  jednego koloru
