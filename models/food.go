package models

import "fmt"

type Dish struct {
	Name      string
	NumMorsel int
}

type Meal struct {
	Dishes []Dish
}

func (m *Meal) String() string {
	meal := fmt.Sprintf("The meal has ")
	for _, val := range m.Dishes {
		meal += fmt.Sprintf("%d morsels of %s ", val.NumMorsel, val.Name)
	}
	return meal
}
