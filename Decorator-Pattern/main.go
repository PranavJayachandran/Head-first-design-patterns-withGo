package main

import (
	"fmt"
	"reflect"
)

// IBeverage interface and its methods
type IBeverage interface {
	getPrice() float32
	addAddOn(string) IBeverage
}

// Black concrete decorator
type Black struct {
	beverage IBeverage
}

func (b *Black) getPrice() float32 {
	return 0.1 * b.beverage.getPrice()
}

func (b *Black) addAddOn(string) IBeverage {
	return b
}

// Sugar concrete decorator
type Sugar struct {
	beverage IBeverage
}

func (s *Sugar) getPrice() float32 {
	return 0.1 * s.beverage.getPrice()
}

func (s *Sugar) addAddOn(string) IBeverage {
	return s
}

// Tea concrete beverage
type Tea struct{}

func (t *Tea) getPrice() float32 {
	return 0.3
}

func (t *Tea) addAddOn(addOn string) IBeverage {
	switch addOn {
	case "Ginger":
		return &Ginger{tea: t}
	default:
		return t
	}
}

// Ginger concrete decorator for Tea
type Ginger struct {
	tea *Tea
}

func (g *Ginger) getPrice() float32 {
	return 0.4 * g.tea.getPrice()
}

func (g *Ginger) addAddOn(string) IBeverage {
	return g
}

// Coffee concrete beverage
type Coffee struct{}

func (c *Coffee) getPrice() float32 {
	return 0.4
}

func (c *Coffee) addAddOn(addOn string) IBeverage {
	switch addOn {
	case "Mocha":
		return &Mocha{coffee: c}
	default:
		return c
	}
}

// Mocha concrete decorator for Coffee
type Mocha struct {
	coffee *Coffee
}

func (m *Mocha) getPrice() float32 {
	return 0.1 * m.coffee.getPrice()
}

func (m *Mocha) addAddOn(string) IBeverage {
	return m
}

// Get the list of possible add-ons based on beverage type
func getAddOnList(b IBeverage) []string {
	switch reflect.TypeOf(b) {
	case reflect.TypeOf(&Tea{}):
		return []string{"Ginger"}
	case reflect.TypeOf(&Coffee{}):
		return []string{"Mocha"}
	default:
		return nil
	}
}

// Add a decorator to the beverage based on the add-on choice
func addBeverageAddOn(addOn string, beverage IBeverage) IBeverage {
	switch addOn {
	case "Black":
		return &Black{beverage: beverage}
	case "Sugar":
		return &Sugar{beverage: beverage}
	default:
		return beverage
	}
}

func main() {
	var beverage IBeverage
	fmt.Println("Enter the choice of Beverage\n1.Tea\n2.Coffee")
	beverageChoice := 0
	fmt.Scanf("%d", &beverageChoice)
	switch beverageChoice {
	case 1:
		beverage = &Tea{}
	case 2:
		beverage = &Coffee{}
	default:
		fmt.Println("Not a valid beverage")
		return
	}

	beverageAddOns := []string{"Black", "Sugar"}
	var addOns int
	fmt.Println("Do you want Addons? (1 for Yes, 0 for No)")
	fmt.Scanf("%d", &addOns)
	for addOns == 1 {
		addOnList := getAddOnList(beverage)
		count := 1
		for _, element := range beverageAddOns {
			fmt.Printf("%d.%s\n", count, element)
			count++
		}
		for _, element := range addOnList {
			fmt.Printf("%d.%s\n", count, element)
			count++
		}
		fmt.Printf("%d.No more addons\n", count)

		fmt.Scanf("%d", &addOns)
		if addOns <= len(beverageAddOns) {
			beverage = addBeverageAddOn(beverageAddOns[addOns-1], beverage)
		} else if addOns <= len(beverageAddOns)+len(addOnList) {
			beverage = beverage.addAddOn(addOnList[addOns-len(beverageAddOns)-1])
		} else {
			break
		}
		addOns = 1
		fmt.Printf("Current price: %.2f\n", beverage.getPrice())
	}
}

// Here the hierarchy goes as follows, tea and coffee are beverages. Sugar and Black are decorators common for
// Tea and Coffee hence they expect the incoming object to be type beverage. Ginger is a decorator for tea
// and mocha is for coffee.
