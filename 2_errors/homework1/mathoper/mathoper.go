package mathoper

type MathOperationMethod interface {
	Calculate(a, b float64) (float64, error)
}

type MathOperationModule struct {
	mathOperationMethod MathOperationMethod
}

func NewMathOperationModule(mathOperationMethod MathOperationMethod) *MathOperationModule {
	return &MathOperationModule{
		mathOperationMethod: mathOperationMethod,
	}
}

// Метод Calculate()
// Принимает:
// 1. Описание проводимой операции
// 2. Значения, над которыми будет проводиться операция
// Возвращает:
// 1. Значение проведённой операции и ошибку
func (m *MathOperationModule) Calculate(a, b float64) (float64, error) {
	if m == nil || m.mathOperationMethod == nil {
		return 0, nil
	}
	return m.mathOperationMethod.Calculate(a, b)
}
