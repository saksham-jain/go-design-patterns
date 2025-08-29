package in_memory 

import (
	"go-design-patterns/strategy_pattern/cache/eviction"
	"go-design-patterns/strategy_pattern/cache"
)

type InMemory struct {
	capacity int64
	evictionStrategy eviction.Eviction 
}

func NewInMemoryCache(capacity int64, evictionStrategy eviction.Eviction) cache.Cache {
		return &InMemory{
			capacity: capacity,
			evictionStrategy: evictionStrategy,
		}
}

func (m *InMemory) Get() {

}

func (m *InMemory) Set() {
	
}
