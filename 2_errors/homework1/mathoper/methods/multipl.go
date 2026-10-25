package methods

// Умножение
type Multiplication struct{}

func NewMultiplication() Multiplication {
	return Multiplication{}
}

func (c Multiplication) Calculate(a, b float64) (float64, error) {
	if a < -1000 || a > 1000 || b < -1000 || b > 1000 {
		return 0, ErrArgOutOfRange
	}
	return a * b, nil
}
