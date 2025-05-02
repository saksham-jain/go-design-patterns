package adapterpattern

import "fmt"

type ThreePinCharger struct {
	name string
	watt string
}

func NewThreePinCharger(name, watt string) *ThreePinCharger {
	return &ThreePinCharger{
		name: name,
		watt: watt,
	}
}

func (t *ThreePinCharger) Charge() error {
	fmt.Printf("charging using three pin charger, charger details: %+v \n", t)
	return nil
}
