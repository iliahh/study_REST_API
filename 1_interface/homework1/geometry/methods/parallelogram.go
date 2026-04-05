package methods

type Parallelogram struct {
	Base   float64
	Side   float64
	Height float64
}

func NewParallelogram(base, side, height float64) Parallelogram {
	return Parallelogram{
		Base:   base,
		Side:   side,
		Height: height,
	}
}

func (p Parallelogram) Area() float64 {
	return p.Base * p.Height
}

func (p Parallelogram) Perimeter() float64 {
	return 2 * (p.Base + p.Side)
}

func (p Parallelogram) Name() string {
	return "Параллелограмм"
}
