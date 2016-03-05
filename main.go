package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pesedr/concurrency-exercises/system"
)

func main() {
	rand.Seed(time.Now().Unix())
	randomNumber := system.RandomNumberGenerator(60, 90)
	fmt.Println("Hello World! ", randomNumber)
}
