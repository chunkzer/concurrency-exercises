package system

import "math/rand"

func RandomNumberGenerator(min, max int) int {
	return min + rand.Intn(max-min)
}
