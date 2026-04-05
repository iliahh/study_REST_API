package methods

type Rhombus struct {
	Side   float64
	Height float64
}

func NewRhombus(side, height float64) Rhombus {
	return Rhombus{
		Side:   side,
		Height: height,
	}
}

func (r Rhombus) Area() float64 {
	return r.Side * r.Height
}

func (r Rhombus) Perimeter() float64 {
	return 4 * r.Side
}

func (r Rhombus) Name() string {
	return "Ромб"
}
