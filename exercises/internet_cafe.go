package exercises

import (
	"fmt"

	"github.com/pesedr/concurrency-exercises/models"
	"github.com/pesedr/concurrency-exercises/system"
)

func Cafe() {
	var tourists []int

	for i := 0; i < 25; i++ {
		tourists = append(tourists, i+1)
	}

	q := models.NewQueue(25)
	for i := 25; i > 0; i-- {
		ran := system.RandomNumberGenerator(0, i)
		q.Push(&models.Node{tourists[ran]})
		fmt.Println(tourists[ran])
		tourists[ran] = tourists[len(tourists)-1]
		tourists = tourists[:len(tourists)-1]
	}

	useComputer()
}

func useComputer() {
	minutes := system.Sleep(15, 120)
	fmt.Println("Tourist ", minutes)
}

type Computer struct {
	Occupied chan bool
	Time     int
}
