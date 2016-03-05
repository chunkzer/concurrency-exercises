package exercises

import "fmt"

func GoForWalk() {
	var alice = &Human{Name: "Alice"}
	var bob = &Human{Name: "Bob"}

	done := make(chan bool)
	alarm := make(chan bool)
	go bob.GetReady(done)
	go alice.GetReady(done)
	go armAlarm(<-done && <-done, alarm)
	go bob.PutOnShoes(done)
	go alice.PutOnShoes(done)
	if <-done && <-done {
		fmt.Println("Exiting and locking the door")
	}
	if <-alarm {
		fmt.Println("Alarm is armed")
	}

}

func armAlarm(done bool, alarm chan bool) {
	fmt.Println("Arming alarm")
	sleep(2, 5)
	fmt.Println("Alarm is counting down")
	sleep(59, 60)
	alarm <- true
}
