package main

import (
	"math/rand"
	"time"

	"github.com/pesedr/concurrency-exercises/exercises"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	exercises.GoForWalk()
	// exercises.EatMeal()
}
