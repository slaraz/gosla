package v2

import (
	"fmt"
	"testing"
	"time"
)

func Test_NowyExchanger_PublikujJSON(t *testing.T) {
	MOJ_EX := `---
url: amqp://guest:guest@localhost:5672/krolik
exchanger:
  nazwa: moj testowy
  kind: fanout
  rodzaj: std
`

	mojex := MusiNowyExchanger(MOJ_EX)
	defer mojex.Close()

	a := struct {
		Id   int
		Imie string
	}{5, "ala"}

	if err := mojex.PublikujJSON(a); err != nil {
		t.Errorf("błąd PublikujJSON(): %v", err)
	}
	err := mojex.Close()
	time.Sleep(100 * time.Millisecond)
	fmt.Println(err)
}
