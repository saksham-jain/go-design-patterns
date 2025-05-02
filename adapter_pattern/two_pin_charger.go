package adapterpattern

import "fmt"

type TwoPinCharger struct {
	name string
	watt string
}

func NewTwoPinCharger(name, watt string) *TwoPinCharger {
	return &TwoPinCharger{
		name: name,
		watt: watt,
	}
}

func (t *TwoPinCharger) Charge() error {
	fmt.Printf("charging using two pin charger, charger details: %+v \n", t)
	return nil
}
