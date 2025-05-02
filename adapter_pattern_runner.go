package main

import (
	adapter "go-design-patterns/adapter_pattern"
)

func main() {
	charger1 := adapter.NewChargerAdapter(adapter.TwoPin, "sj_2_pin", "16Watt")
	charger1.Charge()

	charger2 := adapter.NewChargerAdapter(adapter.ThreePin, "sj_3_pin", "32Watt")
	charger2.Charge()
}
