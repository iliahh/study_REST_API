package methods

type Trapezoid struct {
	BaseA  float64
	BaseB  float64
	SideC  float64
	SideD  float64
	Height float64
}

func NewTrapezoid(baseA, baseB, sideC,
	sideD, height float64) Trapezoid {
	return Trapezoid{
		BaseA:  baseA,
		BaseB:  baseB,
		SideC:  sideC,
		SideD:  sideD,
		Height: height,
	}
}

func (t Trapezoid) Area() float64 {
	return 0.5 * (t.BaseA + t.BaseB) * t.Height
}

func (t Trapezoid) Perimeter() float64 {
	return t.BaseA + t.BaseB + t.SideC + t.SideD
}

func (t Trapezoid) Name() string {
	return "Трапеция"
}
