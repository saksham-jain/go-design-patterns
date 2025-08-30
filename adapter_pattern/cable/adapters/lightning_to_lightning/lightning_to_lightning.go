package lightningtolightning

import (
	"go-design-patterns/adapter_pattern/cable"
)

type LightningToLightningAdapter struct {
	computer cable.Mac // have a concrete class
}

func NewLightningToLightningAdapter(computer cable.Mac) cable.Adapter {
	return &LightningToLightningAdapter{
		computer: computer,
	}
}

func (l *LightningToLightningAdapter) InsertToLightningPort() {
	l.computer.PlugInLightning()
}
