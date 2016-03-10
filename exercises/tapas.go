package exercises

import (
	"fmt"

	"github.com/pesedr/concurrency-exercises/system"
)

func makeMeal() {}

func EatMeal() {
	mealChan := make(chan Meal)

	meal := CreateMeal()
	mealChan <- meal

	alice := Human{Name: "Alice"}
	bob := Human{Name: "Bob"}
	chancla := Human{Name: "Chancla"}
	delta := Human{Name: "Delta"}

	go alice.Choose(mealChan)
	go bob.Choose(mealChan)
	go chancla.Choose(mealChan)
	go delta.Choose(mealChan)
	<-alice.IsReady
	fmt.Println("alice is done")
}

func CreateMeal() Meal {
	var dishes [5]Dish
	var meal Meal
	for i := 0; i < 5; i++ {
		numOfMorsels := system.RandomNumberGenerator(5, 10)
		dishes[i].NumMorsel = numOfMorsels
	}
	meal.Dishes = dishes
	return meal
}
