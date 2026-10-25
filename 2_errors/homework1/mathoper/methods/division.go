package methods

import "errors"

var ErrDivisionByZero = errors.New("деление на ноль")

// Проверка диапазона значений

// Деление
type Division struct{}

func NewDivision() Division {
	return Division{}
}

func (c Division) Calculate(a, b float64) (float64, error) {
	if a < -1000 || a > 1000 || b < -1000 || b > 1000 {
		return 0, ErrArgOutOfRange
	}
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}
