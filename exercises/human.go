package exercises

import (
	"fmt"
	"time"

	"github.com/pesedr/concurrency-exercises/system"
)

type FamilyMember interface {
	GetReady()
	ArmAlarm()
	PutOnShoes()
}

type Human struct {
	Name string
}

func (h *Human) GetReady(isReady chan bool) {
	fmt.Printf("%s is getting ready \n", h)
	seconds := sleep(60, 90)
	fmt.Printf("%s took %d getting ready \n", h, seconds)
	isReady <- true
}

func (h *Human) PutOnShoes(hasShoes chan bool) {
	fmt.Printf("%s is putting on shoes \n", h)
	seconds := sleep(35, 45)
	fmt.Printf("%s took %d seconds putting on shoes \n", h, seconds)
	hasShoes <- true
}

func sleep(min, max int) int {
	randomNumber := system.RandomNumberGenerator(min, max)
	time.Sleep(time.Duration(randomNumber))
	return randomNumber
}

func (h *Human) String() string {
	return h.Name
}
