package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/slaraz/gosla/usuwaczlini/usuwacz"
)

func main() {
	nazwapliku := "test.txt"
	plik, err := os.Open(nazwapliku)
	if err != nil {
		log.Fatal("błąd otwierania pliku", err)
	}
	pokapoka(plik)
}

func pokapoka(r io.Reader) {

	rdr := bufio.NewReader(usuwacz.NowyUsuwacz(r))
	for {
		c, s, err := rdr.ReadRune()
		if err == io.EOF {
			fmt.Println("koniec pliku.")
			break
		} else if err != nil {
			log.Fatal("błąd odczytu runy", err, s)
			break
		}
		fmt.Println(c, strconv.QuoteRune(c))
	}
}
