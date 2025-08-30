package useriterator

import (
	"go-design-patterns/iterator_pattern/iterator"
)

type UserIterator struct {
	list  []string
	index int
}

func NewUserIterator(collection []string) iterator.Iterator {
	return &UserIterator{
		list:  collection,
		index: 0,
	}
}

func (u *UserIterator) HasNext() bool {
	return len(u.list) > u.index
}

func (u *UserIterator) GetNext() string {
	var nextElement string
	if u.HasNext() {
		nextElement = u.list[u.index]
		u.index += 1
	}
	return nextElement
}
