package shape

import "fmt"

type Square struct {
	l float64
	b float64
}

func NewSquare(l, b float64) Shape {
	return &Square{
		l: l,
		b: b,
	}
}

func (s *Square) Area() float64 {
	area := s.l * s.b
	fmt.Printf("square area: %f\n", area)
	return area
}
