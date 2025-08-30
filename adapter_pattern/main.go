package main

import (
	"go-design-patterns/adapter_pattern/cable"
	lightningtolightning "go-design-patterns/adapter_pattern/cable/adapters/lightning_to_lightning"
	lightningtousb "go-design-patterns/adapter_pattern/cable/adapters/lightning_to_usb"
	"go-design-patterns/adapter_pattern/cable/lightning_cable"
)

func main() {
	cable1 := lightning_cable.NewLightningCable("Lightning Cable")

	mac := cable.NewMac("MacBook Pro")
	windows := cable.NewWindows("Dell XPS")

	lightningToLightningAdapter := lightningtolightning.NewLightningToLightningAdapter(mac)
	cable1.InsertToComputer(lightningToLightningAdapter)

	lightningToUSBAdapter := lightningtousb.NewLightningToUsbAdapter(windows)
	cable1.InsertToComputer(lightningToUSBAdapter)
}
