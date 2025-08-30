package cable

import "fmt"

type Windows struct {
	name string
}

func NewWindows(name string) Windows {
	return Windows{
		name: name,
	}
}

func (w *Windows) PlugInUSB() {
	fmt.Printf("plugged in USB to %s\n", w.name)
}
