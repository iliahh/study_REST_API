package methods

import "math"

type Circle struct {
	Radius float64
}

func NewCircle(radius float64) Circle {
	return Circle{
		Radius: radius,
	}
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Name() string {
	return "Круг"
}
