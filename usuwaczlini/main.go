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
	pokapoka(usuwacz.NowyReader(plik))
}

func pokapoka(r io.Reader) {

	rdr := bufio.NewReader(r)
	for {
		c, s, err := rdr.ReadRune()
		if err == io.EOF {
			fmt.Println("koniec pliku.")
			break
		} else if err != nil {
			log.Fatal("błąd odczytu runy: ", err, s)
			break
		}
		fmt.Println(c, strconv.QuoteRune(c))
	}
}

// func skaner() {
// 	var out bytes.Buffer
// 	scanner := bufio.NewScanner(stdout)
// 	for scanner.Scan() {
// 		p := scanner.Bytes()
// 		lines := bytes.Split(p, cr)
// 		out.Write(lines[0])
// 		out.Write(lf)
// 		if len(lines) > 1 {
// 			out.Write(lines[len(lines)-1])
// 			out.Write(lf)
// 		}
// 	}
// }
