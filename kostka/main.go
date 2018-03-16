package main

import (
	"fmt"
	"math/rand"

	"github.com/slaraz/gosla/kostka/kostka2"
)

func losowyRuch(r []func()) {
	x := rand.Intn(len(r))
	r[x]()
}

func mieszaj(r []func(), n int) {
	for i := 0; i < n; i++ {
		losowyRuch(r)
	}
}

func main() {
	fmt.Println("Start...")

	k := kostka2.NowaKostka()
	r := k.WszystkieRuchy()
	mieszaj(r, 3)
	k.Drukuj()

	// 	hist := make(map[kostka2.Kostka]struct{})

	// 	for i := 0; i < 5; i++ {
	// 		losowyRuch(r)
	// 		hist[*k] = struct{}{}
	// 	}
	// 	k.Drukuj()
	// 	fmt.Println("długość hist:", len(hist))

	// 	printMem()

	// 	i := 0
	// 	tout := time.After(1 * time.Second)
	// loop:
	// 	for {

	// 		losowyRuch(r)
	// 		i++

	// 		if i%1000 == 0 {
	// 			select {
	// 			case <-tout:
	// 				break loop
	// 			default:
	// 			}
	// 		}
	// 	}

	// 	fmt.Println("1 sek.:", i, "ruchów")

}
