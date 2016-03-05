package exercises

import (
	"fmt"
	"time"

	"github.com/pesedr/concurrency-exercises/system"
)

type Human interface {
	GetReady()
	ArmAlarm()
	PutOnShoes()
}

type human struct {
	Name string
}

var Person Human

func init() {
	Person = &human{}
}

func (h *human) GetReady() {
	fmt.Println("%s is getting ready", h)
	seconds := sleep(60, 90)
	fmt.Println("%s took %d getting ready", h, seconds)
}

func (h *human) ArmAlarm() {
	fmt.Println("%s is arming the alarm", h)
	sleep(2, 5)
	fmt.Println("Alarm is armed")
}

func (h *human) PutOnShoes() {
	fmt.Println("%s is putting on shoes", h)
	seconds := sleep(35, 45)
	fmt.Println("%s took %d seconds putting on shoes", h, seconds)
}

func sleep(min, max int) int {
	randomNumber := system.RandomNumberGenerator(min, max)
	time.Sleep(time.Duration(randomNumber))
	return randomNumber
}

func (h *human) String() string {
	return h.Name
}
