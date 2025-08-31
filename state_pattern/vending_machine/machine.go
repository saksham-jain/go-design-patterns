package vendingmachine

type Machine interface {
	InsertCoin() string
	DispenseItem() string
	ChangeState(state State)
}
