// +build ignore
package main

import (
	"html/template"
	"log"
	"os"
)

var plik = "ruchy.go"

var dane = []struct {
	Nazwa  string
	Obroty [][]int
}{
	{"ObrótXGóraDół", [][]int{
		{0, 5, 7, 2},
		{3, 6, 4, 1},
		{8, 16, 24, 47},
		{9, 17, 25, 46},
		{10, 18, 26, 45},
	}},
}

var szablon = `
// Code generated automatiko; DO NOT EDIT.
package kostka2

{{- range .}}
func (k *Kostka) {{.Nazwa}}() {
	{{- range .Obroty}}
	buf = k.Kwadraciki[{{index . 0}}]
	k.Kwadraciki[{{index . 0}}] = k.Kwadraciki[{{index . 1}}]
	k.Kwadraciki[{{index . 1}}] = k.Kwadraciki[{{index . 2}}]
	k.Kwadraciki[{{index . 2}}] = k.Kwadraciki[{{index . 3}}]
	k.Kwadraciki[{{index . 3}}] = buf

	{{- end}}
}

{{- end }}
`

func main() {
	t, err := template.New("").Parse(szablon)
	if err != nil {
		log.Fatal("błąd templata:", err)
	}
	f, err := os.Create(plik)
	if err != nil {
		log.Fatal("błąd pliku", err)
	}
	defer f.Close()
	err = t.Execute(f, dane)
	if err != nil {
		log.Fatal("błąd generowania:", err)
	}
}
