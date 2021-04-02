package v2

import (
	"fmt"
	"testing"
	"time"
)

func Test_NowyExchanger_PublikujJSON(t *testing.T) {
	// Go
	RABBIT := "amqp://guest:guest@localhost:5672/krolik"

	konf := konfiguracjaEx{
		url: RABBIT, 
		nazwa: "moj testowy", 
		kind: "fanout", 
		rodzaj: "std", 
		routingKey: "",
	}
	
	// YAML
	konfYAML := `---
url: amqp://guest:guest@localhost:5672/krolik
exchanger:
  nazwa: moj testowy
  kind: fanout
  rodzaj: std
` // niestety trzeba pilnować wcięć

	// JSON
	konfJSON := `{
		"url": "amqp://guest:guest@localhost:5672/krolik",
		"exchanger": {
			"nazwa": "moj testowy",
			  "kind": "fanout",
			  "rodzaj": "std"
		}
	}`
	
	konfJSONSimple := `{"nazwa": "moj testowy", "kind": "fanout", "rodzaj": "std"}`

	konfTOML := `
[exchanger]	
nazwa = "moj testowy"
kind = "fanout"
rodzaj = "std"
`

	konfSla1 := "exchange=moj testowy; kind=fanout; rodzaj=std"
	konfSla2 := "url=amqp://guest:guest@localhost:5672/krolik;exchange=moj testowy; kind=fanout; rodzaj=std"


	konfSla3 := "exchange=moj testowy; url=amqp://guest:guest@localhost:5672/krolik; kind=fanout; rodzaj=std;"
	
	mojex := MusiNowyExchanger(RABBIT, "exchange=moj testowy; kind=fanout; rodzaj=std")
	mojex := MusiNowyExchanger(RABBIT, "moj testowy", "fanout", "std")
	defer mojex.Close()

	a := struct {
		Id   int
		Imie string
	}{5, "ala"}

	ladna(/*a */15, b:="alal")

	if err := mojex.PublikujJSON(a); err != nil {
		t.Errorf("błąd PublikujJSON(): %v", err)
	}
	err := mojex.Close()
	time.Sleep(time.Second)
	fmt.Println(err)
}

func ladna(a int, b string) {}