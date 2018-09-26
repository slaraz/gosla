package usuwacz

import "io"

type Usuwacz struct {
	r io.Reader
}

func (u *Usuwacz) Read(b []byte) (n int, err error) {
	return 0, nil
}
