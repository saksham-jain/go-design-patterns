package states

import vendingmachine "go-design-patterns/state_pattern/vending_machine"

type PaymentDoneState struct {
	context vendingmachine.Machine
}

func NewPaymentDoneState(vendingMachine vendingmachine.Machine) vendingmachine.State {
	return &PaymentDoneState{
		context: vendingMachine,
	}
}

func (p *PaymentDoneState) SetContext(vendingMachine vendingmachine.Machine) {
	p.context = vendingMachine
}

func (p *PaymentDoneState) InsertCoin() string {
	return "Coin already inserted, Please select an item"
}

func (p *PaymentDoneState) DispenseItem() string {
	p.context.ChangeState(NewOpenForOrderState(p.context))
	return "Item dispensed, Please insert coin for next order"
}
