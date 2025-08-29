package lru

import "go-design-patterns/strategy_pattern/cache/eviction"

type Lru struct {

}

func NewLu() eviction.Eviction {
	return &Lru{}
}

func (l *Lru) Evict() {

}
