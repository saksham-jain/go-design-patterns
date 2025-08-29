package main

import (
	"fmt"
	"go-design-patterns/strategy_pattern/cache/eviction/lru"
	"go-design-patterns/strategy_pattern/cache/in_memory"
)

// example of in-memory cache with different eviction policy

func main() {
	evictionPolicyObj := lru.NewLu()
	cacheObj := in_memory.NewInMemoryCache(2, evictionPolicyObj)
	fmt.Printf("%+v", cacheObj)
}
