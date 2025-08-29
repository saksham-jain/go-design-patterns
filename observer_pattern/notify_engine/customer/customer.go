package customer

import (
	"fmt"
	"go-design-patterns/observer_pattern/notify_engine"
)

type Customer struct {
	name string
}

func NewCustomer(name string) notify_engine.Observer {
	return &Customer{
		name: name,
	}
}

func (c *Customer) Notify() {
	fmt.Printf("customer %s is notified\n", c.name)
}
