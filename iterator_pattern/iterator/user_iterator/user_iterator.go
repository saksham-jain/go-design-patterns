package useriterator

import (
	"go-design-patterns/iterator_pattern/iterator"
	usercollection "go-design-patterns/iterator_pattern/iterator/user_collection"
)

type UserIterator struct {
	list  usercollection.UserCollection
	index int
}

func NewUserIterator(collection usercollection.UserCollection) iterator.Iterator {
	return &UserIterator{
		list:  collection,
		index: 0,
	}
}

func (u *UserIterator) HasNext() bool {
	return len(u.list.Collection) > u.index
}

func (u *UserIterator) GetNext() string {
	var nextElement string
	if u.HasNext() {
		nextElement = u.list.Collection[u.index]
		u.index += 1
	}
	return nextElement
}
