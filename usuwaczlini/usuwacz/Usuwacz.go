package usuwacz

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
)

type Usuwacz struct {
	rdr io.Reader
}

func NowyUsuwacz(r io.Reader) *Usuwacz {
	return &Usuwacz{rdr: r}
}

func (u *Usuwacz) Read(b []byte) (n int, err error) {
	buf, err := ioutil.ReadAll(u.rdr)
	bbuf := bytes.NewBuffer(buf)
	i := 0
	for i := range b {
		c, err := bbuf.ReadByte()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("błąd czytania strumienia:", err)
		}
		b[i] = c

	}
	return i, err
}
