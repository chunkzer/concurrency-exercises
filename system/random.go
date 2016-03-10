package system

import (
	"math/rand"
	"time"
)

func RandomNumberGenerator(min, max int) int {
	return min + rand.Intn(max-min)
}

func Sleep(min, max int) int {
	randomNumber := RandomNumberGenerator(min, max)
	time.Sleep(time.Duration(randomNumber) * time.Millisecond)
	return randomNumber
}
