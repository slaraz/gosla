package poszukiwanie

import (
	"fmt"
	"os"
	"time"

	"github.com/slaraz/gosla/kostka/kostka2"
)

func Szukaj(k *kostka2.Kostka) {
	tout := time.After(10 * time.Second)
szukanie:
	for {
		fmt.Println("szukam...")
		czyUlozona(k)
		kombinacje := kombinuj(k)
		k = najlepsza(kombinacje)

		// timeout
		select {
		case <-tout:
			break szukanie
		default:
		}
	}
	k.Drukuj()
}

var ulozona = kostka2.NowaKostka()

func czyUlozona(k *kostka2.Kostka) {
	if *k == *ulozona {
		fmt.Println("Ułożona!!!")
		os.Exit(0)
	}
}

func najlepsza(kk []kostka2.Kostka) *kostka2.Kostka {
	naj := 0
	najk := kk[0]
	for _, k := range kk {
		o := k.IleNaMiejscu()
		if naj < o {
			naj = o
			najk = k
		}
	}
	return &najk
}

func ocena(k kostka2.Kostka) int {
	return 5
}
