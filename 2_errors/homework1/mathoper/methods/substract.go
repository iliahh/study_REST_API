package methods

// Вычитание
type Subtraction struct{}

func NewSubtraction() Subtraction {
	return Subtraction{}
}

func (c Subtraction) Calculate(a, b float64) (float64, error) {
	if a < -1000 || a > 1000 || b < -1000 || b > 1000 {
		return 0, ErrArgOutOfRange
	}
	return a - b, nil
}
