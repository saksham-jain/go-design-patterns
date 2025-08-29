package main

import (
	"go-design-patterns/observer_pattern/notify_engine/customer"
	"go-design-patterns/observer_pattern/notify_engine/item"
)

func main() {
	customer1 := customer.NewCustomer("Ramesh")
	customer2 := customer.NewCustomer("Suresh")

	item1 := item.NewItem("Pen")
	item1.Register(customer1)
	item1.Register(customer2)
	item1.UpdateInventory(2)
}
