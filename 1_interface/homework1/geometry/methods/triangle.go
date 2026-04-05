package methods

type Triangle struct {
	Base   float64
	SideA  float64
	SideB  float64
	Height float64
}

func NewTriangle(base, height, sideA, sideB float64) Triangle {
	return Triangle{
		Base:   base,
		SideA:  sideA,
		SideB:  sideB,
		Height: height,
	}
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func (t Triangle) Perimeter() float64 {
	return t.Base + t.SideA + t.SideB
}

func (t Triangle) Name() string {
	return "Треугольник"
}
