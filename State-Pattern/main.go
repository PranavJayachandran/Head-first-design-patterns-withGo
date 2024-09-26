package main

import "fmt"

type State interface {
	addItem(int)
	requestItem()
	pay(int)
	dispense()
}

type vendor struct {
	hasItem       State
	itemRequested State
	noItem        State
	hasMoney      State

	currentState State
	itemCount    int
}

func (v *vendor) requestItem() {
	v.currentState.requestItem()
}

func (v *vendor) addItem(count int) {
	v.currentState.addItem(count)
}

func (v *vendor) insertMoney(money int) {
	v.currentState.pay(money)
}

func (v *vendor) dispenseItem() {
	v.currentState.dispense()
}

// Has item State
type hasItemState struct {
	v *vendor
}

func (h *hasItemState) addItem(count int) {
	h.v.itemCount += count
}
func (h *hasItemState) requestItem() {
	h.v.itemCount--
	fmt.Println("Item requested has been found")
	h.v.currentState = h.v.itemRequested
}
func (h *hasItemState) pay(pay int) {
	fmt.Println("Request item first")
}
func (h *hasItemState) dispense() {
	fmt.Println("Request item first")
}

// Item Reqested State
type ItemRequestedState struct {
	v *vendor
}

func (h *ItemRequestedState) addItem(count int) {
	h.v.itemCount += count
}
func (h *ItemRequestedState) requestItem() {
	fmt.Println("Item has been requested pay the money")
}
func (h *ItemRequestedState) pay(pay int) {
	fmt.Println("Money has been paid")
	h.v.currentState = h.v.hasMoney
}
func (h *ItemRequestedState) dispense() {
	fmt.Println("Pay for the item first")
}

// No Item State
type NoItemState struct {
	v *vendor
}

func (h *NoItemState) addItem(count int) {
	h.v.itemCount += count
	h.v.currentState = h.v.hasItem
}
func (h *NoItemState) requestItem() {
	fmt.Println("No Item available")
}
func (h *NoItemState) pay(pay int) {
	fmt.Println("Request item first")
}
func (h *NoItemState) dispense() {
	fmt.Println("Request item first")
}

// Has Money State
type HasMoneyState struct {
	v *vendor
}

func (h *HasMoneyState) addItem(count int) {
	h.v.itemCount += count
}
func (h *HasMoneyState) requestItem() {
	fmt.Println("Dispese in process")
}
func (h *HasMoneyState) pay(pay int) {
	fmt.Print("Dispense in peocess")
}
func (h *HasMoneyState) dispense() {
	fmt.Println("Item dispensed")
	if h.v.itemCount == 0 {
		h.v.currentState = h.v.noItem
	} else {
		h.v.currentState = h.v.hasItem
	}
}

func main() {
	v := &vendor{
		itemCount: 1,
	}
	hasItemState := &hasItemState{
		v: v,
	}
	itemRequestedState := &ItemRequestedState{
		v: v,
	}
	hasMoneyState := &HasMoneyState{
		v: v,
	}
	noItemState := &NoItemState{
		v: v,
	}
	v.currentState = hasItemState
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState

	v.requestItem()
	v.insertMoney(10)
	v.dispenseItem()

	v.requestItem()
	v.addItem(1)
	v.requestItem()
}
