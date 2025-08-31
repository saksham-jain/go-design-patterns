package states

import vendingmachine "go-design-patterns/state_pattern/vending_machine"

type OpenForOrderState struct {
	context vendingmachine.Machine
}

func NewOpenForOrderState(vendingMachine vendingmachine.Machine) vendingmachine.State {
	return &OpenForOrderState{
		context: vendingMachine,
	}
}

func (o *OpenForOrderState) SetContext(vendingMachine vendingmachine.Machine) {
	o.context = vendingMachine
}

func (o *OpenForOrderState) InsertCoin() string {
	o.context.ChangeState(NewPaymentDoneState(o.context))
	return "Coin Inserted, You can now select an item"
}

func (o *OpenForOrderState) DispenseItem() string {
	return "Please insert coin first"
}
