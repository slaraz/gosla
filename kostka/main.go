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

type IKostka interface {
	ObrotYGoraLewo() *IKostka
	ObrotYGoraPrawo() *IKostka
	// ObrotYSrodekLewo() *IKostka
	// ObrotYSrodekPrawo() *IKostka
	// ObrotYDolLewo() *IKostka
	// ObrotYDolPrawo() *IKostka

	// ObrotXLewoGora() *IKostka
	// ObrotXLewoDol() *IKostka
	// ObrotXSrodekGora() *IKostka
	// ObrotXSrodekDol() *IKostka
	// ObrotXPrawoGora() *IKostka
	// ObrotXPrawoDol() *IKostka

	// ObrotZPrzodPrawo() *IKostka
	// ObrotZPrzodLewo() *IKostka
	// ObrotZSrodekPrawo() *IKostka
	// ObrotZSrodekLewo() *IKostka
	// ObrotZTylPrawo() *IKostka
	// ObrotZTylLewo() *IKostka
}

type Ruch func() IKostka

var k *IKostka = new(Kostka)

var Ruchy = [](func() *IKostka){
	k.ObrotYGoraLewo(),
	k.ObrotYGoraPrawo(),
}

// 1 1 1
// 1 1 1
// 1 1 1

// 	      1
//      1   1
//    1   1   1
//  2   1   1   3
//  2 2   1   3 3
//  2 2 2   3 3 3
//    2 2   3 3
//      2   3
