package vendingmachine

type State interface {
	SetContext(vendingMachine Machine)
	InsertCoin() string
	DispenseItem() string
}
