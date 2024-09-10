package main

import (
	"fmt"
)

type IPizza interface {
	prepare()
	serve()
}
type VegPizza struct {
}

func (v *VegPizza) prepare() {
	fmt.Println("Preparing Veg Pizza")
}
func (v *VegPizza) serve() {
	fmt.Println("Serving Veg Pizza")
}

type NonVegPizza struct {
}

func (v *NonVegPizza) prepare() {
	fmt.Println("Preparing Non Veg Pizza")
}
func (v *NonVegPizza) serve() {
	fmt.Println("Serving Non Veg Pizza")
}

func getPizza(pizza string) IPizza {
	if pizza == "Veg" {
		return &VegPizza{}
	} else if pizza == "Non-Veg" {
		return &NonVegPizza{}
	}
	return nil
}
func main() {
	var pizzaName string
	fmt.Println("Enter the name of the pizza Veg or Non-Veg")
	fmt.Scanf("%s", &pizzaName)
	pizza := getPizza(pizzaName)
	if pizza == nil {
		fmt.Println("Error")
	} else {
		pizza.prepare()
		pizza.serve()
	}
}
