package exercises

import "fmt"

func GoForWalk() {
	var alice = &Human{Name: "Alice"}
	var bob = &Human{Name: "Bob"}

	one := make(chan bool)
	two := make(chan bool)
	alarm := make(chan bool)
	go bob.GetReady(one)
	go alice.GetReady(two)
	go armAlarm(<-one && <-two, alarm)
	go alarmCountdown(<-alarm, alarm)
	go bob.PutOnShoes(one)
	go alice.PutOnShoes(two)
	<-one
	<-two
	fmt.Println("exiting and locking the door, ")
	<-alarm
}

func alarmCountdown(alarm bool, c chan bool) {
	fmt.Println("Alarm is counting down")
	sleep(59, 60)
	fmt.Println("Alarm is armed")
	c <- true
}
func armAlarm(done bool, alarm chan bool) {
	fmt.Println("Arming alarm")
	sleep(2, 5)
	alarm <- true
}

func leave(done bool) {
	fmt.Println("Exiting and locking the door")
}

func armed(alarm bool) {
}
