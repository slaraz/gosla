package poszukiwanie

import "math/rand"

func losowyRuch(r []func()) {
	x := rand.Intn(len(r))
	r[x]()
}

func Mieszaj(r []func(), n int) {
	for i := 0; i < n; i++ {
		losowyRuch(r)
	}
}
