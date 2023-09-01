package atelier3

import (
	"math/rand"
)

func ArrayGenerate(taille int) map[int]int {
	var x = make(map[int]int)

	for i := 0; i < taille; i++ {
		nb := rand.Intn(10000)
		x[i] = nb
	}

	return x
}