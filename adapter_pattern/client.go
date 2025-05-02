package adapterpattern

func NewChargerAdapter(charger_type ChargerType, name, watt string) ChargerAdapter {
	switch charger_type {
	case TwoPin:
		return NewTwoPinCharger(name, watt)
	case ThreePin:
		return NewThreePinCharger(name, watt)
	default:
		panic("charger type is not valid!")
	}
}
