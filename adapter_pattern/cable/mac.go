package cable

import "fmt"

type Mac struct {
	name string
}

func NewMac(name string) Mac {
	return Mac{
		name: name,
	}
}

func (m *Mac) PlugInLightning() {
	fmt.Printf("plugged in Lightning to %s\n", m.name)
}
