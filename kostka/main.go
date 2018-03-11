package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Start.")
	k := nowaKostka()
	b, _ := json.MarshalIndent(k, "", "  ")
	fmt.Printf(string(b))
}

type ruch func()

func rruru() {
	k := nowaKostka()
	ruchy := []ruch{
		k.ObrotYGoraLewo,
		k.ObrotYGoraPrawo,
	}
	fmt.Println(len(ruchy))
}
