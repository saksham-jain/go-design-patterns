package shape

import (
	"fmt"
)

// factory can also be interface in large code. Its fine to use direct method in small code.
// with interface we can have many factories.
func NewShape(shapeType string, unit float64) Shape {
	switch shapeType {
	case "circle":
		return NewCircle(unit)
	case "square":
		return NewSquare(unit, unit)
	default:
		fmt.Print("invalid shape input")
		return nil
	}
}
