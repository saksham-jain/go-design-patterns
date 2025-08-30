package main

import (
	"fmt"
	usercollection "go-design-patterns/iterator_pattern/iterator/user_collection"
)

func main() {
	list := []string{"Rajesh", "Ramesh", "Suresh"}
	userCollection := usercollection.NewUserCollection(list)

	userIterator := userCollection.Iterator()
	fmt.Printf("next value: %s\n", userIterator.GetNext())
	fmt.Printf("next value: %s\n", userIterator.GetNext())
	fmt.Printf("next value: %s\n", userIterator.GetNext())
	fmt.Printf("next value: %s\n", userIterator.GetNext())
	fmt.Printf("next value: %s\n", userIterator.GetNext())
}
