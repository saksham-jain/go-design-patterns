package main

import (
	vendingmachine "go-design-patterns/state_pattern/vending_machine"
	"go-design-patterns/state_pattern/vending_machine/states"
)

func main() {
	vendingMachine1 := vendingmachine.NewVendingMachine("VendingMachine1")
	initialState := states.NewOpenForOrderState(vendingMachine1)
	initialState.SetContext(vendingMachine1)
	vendingMachine1.ChangeState(initialState)
	println(vendingMachine1.InsertCoin())
	println(vendingMachine1.DispenseItem())
	println(vendingMachine1.DispenseItem())
	println(vendingMachine1.InsertCoin())
	println(vendingMachine1.DispenseItem())
	println(vendingMachine1.InsertCoin())
	println(vendingMachine1.DispenseItem())
	println(vendingMachine1.DispenseItem())
}
