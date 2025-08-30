package main

import (
	"fmt"
	usercollection "go-design-patterns/iterator_pattern/iterator/user_collection"
	useriterator "go-design-patterns/iterator_pattern/iterator/user_iterator"
)

func main() {
	list := []string{"Rajesh", "Ramesh", "Suresh"}
	userCollection := usercollection.NewUserCollection(list)

	userIterator := useriterator.NewUserIterator(userCollection)
	fmt.Printf("next value: %s\n", userIterator.GetNext())
	fmt.Printf("next value: %s\n", userIterator.GetNext())
	fmt.Printf("next value: %s\n", userIterator.GetNext())
	fmt.Printf("next value: %s\n", userIterator.GetNext())
	fmt.Printf("next value: %s\n", userIterator.GetNext())
}
