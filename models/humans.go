package models

import (
	"fmt"

	"github.com/pesedr/concurrency-exercises/system"
)

type Human struct {
	Name    string
	IsReady chan bool
}

func (h *Human) GetReady(isReady chan bool) {
	fmt.Printf("%s is getting ready \n", h)
	seconds := system.Sleep(60, 90)
	fmt.Printf("%s took %d getting ready \n", h, seconds)
	isReady <- true
}

func (h *Human) PutOnShoes(hasShoes chan bool) {
	fmt.Printf("%s is putting on shoes \n", h)
	seconds := system.Sleep(35, 45)
	fmt.Printf("%s took %d seconds putting on shoes \n", h, seconds)
	hasShoes <- true
}

func (h *Human) String() string {
	return h.Name
}

func (h *Human) EatMorsel(meal *Meal, i int) {
	meal.Dishes[i].NumMorsel--
	fmt.Printf("%s is enjoying some %s \n", h, meal.Dishes[i].Name)
	if meal.Dishes[i].NumMorsel == 0 {
		fmt.Printf("No more %s left!\n", meal.Dishes[i].Name)
		meal.Dishes[i] = meal.Dishes[len(meal.Dishes)-1]
		meal.Dishes = meal.Dishes[:len(meal.Dishes)-1]
	}
	system.Sleep(30, 180)
}

func (h *Human) Choose(meal *Meal, isMealDone chan bool) {
	if len(meal.Dishes) == 0 {
		isMealDone <- true
		return
	}
	i := system.RandomNumberGenerator(0, len(meal.Dishes))
	h.EatMorsel(meal, i)
	h.Choose(meal, isMealDone)
}
