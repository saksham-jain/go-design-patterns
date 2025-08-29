package item

import (
	"go-design-patterns/observer_pattern/notify_engine"
)

type Item struct {
	name         string
	inventory    int64
	observerList []notify_engine.Observer
}

func NewItem(name string) notify_engine.Subject {
	return &Item{
		name: name,
	}
}

func (i *Item) Register(o notify_engine.Observer) {
	i.observerList = append(i.observerList, o)
}

func (i *Item) Deregister(o notify_engine.Observer) {

}

func (i *Item) UpdateInventory(c int64) {
	i.inventory = c
	i.notifyAll()
}

func (i *Item) notifyAll() {
	for _, o := range i.observerList {
		o.Notify()
	}
}
