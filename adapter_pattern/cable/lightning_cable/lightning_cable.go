package lightning_cable

import "go-design-patterns/adapter_pattern/cable"

type LightningCable struct {
	name string
}

func NewLightningCable(name string) cable.Cable {
	return &LightningCable{
		name: name,
	}
}

func (l *LightningCable) InsertToComputer(adapter cable.Adapter) {
	adapter.InsertToLightningPort()
}
