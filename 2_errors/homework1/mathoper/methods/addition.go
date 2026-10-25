package methods

// Сложение
type Addition struct{}

func NewAddition() Addition {
	return Addition{}
}

func (c Addition) Calculate(a, b float64) (float64, error) {
	if a < -1000 || a > 1000 || b < -1000 || b > 1000 {
		return 0, ErrArgOutOfRange
	}
	return a + b, nil
}
