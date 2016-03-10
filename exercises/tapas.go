package exercises

import (
	"fmt"

	"github.com/pesedr/concurrency-exercises/system"
)

func makeMeal() {}

func EatMeal() {
	mealChan := make(chan bool)

	meal := createMeal()
	fmt.Printf("%s \n", &meal)

	alice := Human{Name: "Alice"}
	bob := Human{Name: "Bob"}
	chancla := Human{Name: "Chancla"}
	delta := Human{Name: "Delta"}

	go alice.Choose(&meal, mealChan)
	go bob.Choose(&meal, mealChan)
	go chancla.Choose(&meal, mealChan)
	go delta.Choose(&meal, mealChan)
	<-mealChan
	fmt.Println("That was delicious!")
}

func createMeal() Meal {
	var dishes [5]Dish
	var meal Meal

	for i := 0; i < 5; i++ {
		numOfMorsels := system.RandomNumberGenerator(5, 10)
		dishes[i].NumMorsel = numOfMorsels
		meal.Dishes = append(meal.Dishes, dishes[i])
	}

	meal.Dishes[0].Name = "chorizo"
	meal.Dishes[1].Name = "chopitos"
	meal.Dishes[2].Name = "patatas bravas"
	meal.Dishes[3].Name = "pimientos"
	meal.Dishes[4].Name = "croquetas"
	return meal
}
