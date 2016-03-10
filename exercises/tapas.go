package exercises

import (
	"fmt"

	"github.com/pesedr/concurrency-exercises/models"
	"github.com/pesedr/concurrency-exercises/system"
)

func makeMeal() {}

func EatMeal() {
	mealChan := make(chan bool)

	meal := createMeal()
	diners := getDiners()

	for _, diner := range diners {
		go diner.Choose(&meal, mealChan)
	}
	<-mealChan
	fmt.Println("That was delicious!")
}

func getDiners() (diners []*models.Human) {
	diners = append(diners, &models.Human{Name: "Alice"})
	diners = append(diners, &models.Human{Name: "Bob"})
	diners = append(diners, &models.Human{Name: "Chancla"})
	diners = append(diners, &models.Human{Name: "Delta"})
	return diners
}

func createMeal() models.Meal {
	var dishes [5]models.Dish
	var meal models.Meal

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

	fmt.Printf("%s \n", &meal)
	return meal
}
