package main

import (
	"go-design-patterns/factory_pattern/shape"
)

func main() {
	circle := shape.NewShape("circle", 2)
	_ = circle.Area()

	square := shape.NewShape("square", 3)
	_ = square.Area()
}
