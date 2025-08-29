package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func main() {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if singleInstance == nil {
			singleInstance = &single{}
		} else {
			fmt.Print("instance already initialized")
		}
	} else {
		fmt.Print("instance already initialized")
	}

	fmt.Printf("instance: %+v", singleInstance)
}
