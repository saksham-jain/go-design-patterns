package shape

import "fmt"

type Circle struct {
	radius float64
}

func NewCircle(radius float64) Shape {
	return &Circle{
		radius: radius,
	}
}

func (c *Circle) Area() float64 {
	fmt.Print("ha")
	area := 2 * c.radius
	fmt.Printf("circle area: %f\n", area)
	return area
}
