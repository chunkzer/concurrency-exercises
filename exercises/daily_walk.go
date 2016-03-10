package exercises

import (
	"fmt"

	"github.com/pesedr/concurrency-exercises/models"
	"github.com/pesedr/concurrency-exercises/system"
)

func GoForWalk() {
	runners := makeRunners()
	alarm := make(chan bool)
	one := make(chan bool)
	two := make(chan bool)

	go runners[0].GetReady(one)
	go runners[1].GetReady(two)
	go armAlarm(<-one && <-two, alarm)
	go alarmCountdown(<-alarm, alarm)
	go runners[0].PutOnShoes(one)
	go runners[1].PutOnShoes(two)
	<-one
	<-two
	fmt.Println("exiting and locking the door, ")
	<-alarm
}

func makeRunners() (runners []models.Human) {
	runners = append(runners, models.Human{Name: "Alice"})
	runners = append(runners, models.Human{Name: "Bob"})
	return runners
}
func alarmCountdown(alarm bool, c chan bool) {
	fmt.Println("Alarm is counting down")
	system.Sleep(59, 60)
	fmt.Println("Alarm is armed")
	c <- true
}
func armAlarm(done bool, alarm chan bool) {
	fmt.Println("Arming alarm")
	system.Sleep(2, 5)
	alarm <- true
}

func leave(done bool) {
	fmt.Println("Exiting and locking the door")
}

func armed(alarm bool) {
}
