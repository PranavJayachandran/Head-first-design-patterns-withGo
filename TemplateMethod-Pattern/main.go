package main

import "fmt"

// Interfaces and the base class
type Ibeverage interface {
	addIngredients()
	heat()
	pour()
}
type beverage struct {
	beverage Ibeverage
}

func (b *beverage) make() {
	b.beverage.addIngredients()
	b.beverage.heat()
	b.beverage.pour()
}

// Concrete Tea Class
type tea struct{}

func (t *tea) addIngredients() {
	fmt.Println("Adding ingredients to tea")
}
func (t *tea) heat() {
	fmt.Println("Heating tea")
}
func (t *tea) pour() {
	fmt.Println("Pouring tea")
}

// Concrete Coffee Class
type coffee struct{}

func (c *coffee) addIngredients() {
	fmt.Println("Adding ingredients to coffee")
}
func (c *coffee) heat() {
	fmt.Println("Heating coffee")
}
func (c *coffee) pour() {
	fmt.Println("Pouring coffee")
}

func main() {
	tea := &tea{}
	bevearege := &beverage{beverage: tea}
	bevearege.make()

	bevearege2 := &beverage{beverage: &coffee{}}
	bevearege2.make()

}
