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
	{"ObrYGora", [][]int{
		{0, 5, 7, 2},
		{3, 6, 4, 1},
		{8, 16, 24, 47},
		{9, 17, 25, 46},
		{10, 18, 26, 45},
	}},
	{"ObrYDol", [][]int{
		{32, 37, 39, 34},
		{35, 38, 36, 33},
		{13, 21, 29, 42},
		{14, 22, 30, 41},
		{15, 23, 31, 40},
	}},
	{"ObrXLewo", [][]int{
		{0, 16, 32, 40},
		{3, 19, 35, 43},
		{5, 21, 37, 45},
		{8, 13, 15, 10},
		{11, 14, 12, 9},
	}},
	{"ObrXPrawo", [][]int{
		{2, 18, 34, 42},
		{4, 20, 36, 44},
		{7, 23, 39, 47},
		{24, 29, 31, 26},
		{27, 30, 28, 25},
	}},
	{"ObrZPrzod", [][]int{
		{16, 21, 23, 18},
		{19, 22, 20, 17},
		{5, 24, 34, 15},
		{6, 27, 33, 12},
		{7, 29, 32, 10},
	}},
	{"ObrZTyl", [][]int{
		{40, 45, 47, 42},
		{43, 46, 44, 41},
		{37, 31, 2, 8},
		{38, 28, 1, 11},
		{39, 26, 0, 13},
	}},
}

var szablon = `
// Code generated automatiko; DO NOT EDIT.
package kostka2

func (k *Kostka) WszystkieRuchy() []func() {
	return []func() {
	{{- range .}}
		k.{{.Nazwa}}A,
		k.{{.Nazwa}}B,
	{{- end}}
	}
}

{{- range .}}
func (k *Kostka) {{.Nazwa}}A() {
{{- range .Obroty}}
	buf = k.Kwadraciki[{{index . 0}}]
	k.Kwadraciki[{{index . 0}}] = k.Kwadraciki[{{index . 1}}]
	k.Kwadraciki[{{index . 1}}] = k.Kwadraciki[{{index . 2}}]
	k.Kwadraciki[{{index . 2}}] = k.Kwadraciki[{{index . 3}}]
	k.Kwadraciki[{{index . 3}}] = buf
{{- end}}
}
func (k *Kostka) {{.Nazwa}}B() {
{{- range .Obroty}}
	buf = k.Kwadraciki[{{index . 3}}]
	k.Kwadraciki[{{index . 3}}] = k.Kwadraciki[{{index . 2}}]
	k.Kwadraciki[{{index . 2}}] = k.Kwadraciki[{{index . 1}}]
	k.Kwadraciki[{{index . 1}}] = k.Kwadraciki[{{index . 0}}]
	k.Kwadraciki[{{index . 0}}] = buf
{{- end}}
}
{{- end }}
`

func main() {
	t, err := template.New("").Parse(szablon)
	blad("błąd templata:", err)

	f, err := os.Create(plik)
	blad("błąd pliku:", err)
	defer f.Close()

	err = t.Execute(f, dane)
	blad("błąd generowania:", err)
}

func blad(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}
