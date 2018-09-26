package usuwacz

import (
	"io"
)

type Usuwacz struct {
	rdr io.Reader
}

func NowyReader(r io.Reader) *Usuwacz {
	return &Usuwacz{rdr: r}
}

var CR = byte('\r')
var LF = byte('\n')

func (u *Usuwacz) Read(p []byte) (int, error) {
	m, err := u.rdr.Read(p)
	if err != nil {
		return m, err
	}
	buf := make([]byte, m)
	n := 0
	for i := 0; i < m; i++ {
		x := p[i]
		if x == CR || x == LF {
			continue
		}
		buf[n] = x
		n++
	}

	copy(p, buf[:n])
	return n, nil
}
