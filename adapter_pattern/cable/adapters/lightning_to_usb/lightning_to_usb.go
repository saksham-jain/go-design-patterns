// adapter implementation
package lightningtousb

import (
	"fmt"
	"go-design-patterns/adapter_pattern/cable"
)

type LightningToUSBAdapter struct {
	computer cable.Windows // have a concrete class
}

func NewLightningToUsbAdapter(computer cable.Windows) cable.Adapter {
	return &LightningToUSBAdapter{
		computer: computer,
	}
}

func (l *LightningToUSBAdapter) InsertToLightningPort() {
	fmt.Printf("converting lightning to USB..\n")
	l.computer.PlugInUSB()
}
