package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Start.")

	k := nowaKostka()
	ruchy := k.wszystkieRuchy()
	fmt.Println(len(ruchy))

	b, _ := json.MarshalIndent(k, "", "  ")
	fmt.Printf(string(b))
}
