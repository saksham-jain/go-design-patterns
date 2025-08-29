package lfu

import "go-design-patterns/strategy_pattern/cache/eviction"

type Lfu struct {

}

func NewLfu() eviction.Eviction {
	return &Lfu{}
}

func (l *Lfu) Evict() {
	
}
