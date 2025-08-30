package usercollection

import (
	"go-design-patterns/iterator_pattern/iterator"
	useriterator "go-design-patterns/iterator_pattern/iterator/user_iterator"
)

type UserCollection struct {
	Collection []string
}

func NewUserCollection(collection []string) UserCollection {
	return UserCollection{
		Collection: collection,
	}
}

func (u *UserCollection) Iterator() iterator.Iterator {
	return useriterator.NewUserIterator(u.Collection)
}
