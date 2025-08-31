package vendingmachine

type VendingMachine struct {
	name         string
	currentState State
}

func NewVendingMachine(name string) Machine {
	return &VendingMachine{
		name: name,
	}
}

func (v *VendingMachine) InsertCoin() string {
	return v.currentState.InsertCoin()
}

func (v *VendingMachine) DispenseItem() string {
	return v.currentState.DispenseItem()
}

func (v *VendingMachine) ChangeState(state State) {
	v.currentState = state
}
