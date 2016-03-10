package exercises

import (
	"fmt"
	"time"

	"github.com/pesedr/concurrency-exercises/system"
)

type Human struct {
	Name    string
	IsReady chan bool
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

func (h *Human) String() string {
	return h.Name
}

func sleep(min, max int) int {
	randomNumber := system.RandomNumberGenerator(min, max)
	time.Sleep(time.Duration(randomNumber) * time.Millisecond)
	return randomNumber
}

type Morsel string

type Dish struct {
	NumMorsel int
}

type Meal struct {
	Dishes [5]Dish
}

func (h *Human) EatMorsel(morsel Morsel, isDone chan bool) {
	fmt.Printf("%s is putting is enjoying some %s \n", h, morsel)
	sleep(30, 180)
	isDone <- true
}

func (h *Human) Choose(meal chan Meal) {
	gonnaEat <- meal
	i := system.RandomNumberGenerator(0, 4)
	h.EatMorsel(gonnaEat.Dishes[i])

	<-h.IsReady
	h.Choose(meal)
}
