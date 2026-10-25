package main

import (
	"Study_REST_API/2_errors/homework1/mathoper"
	"Study_REST_API/2_errors/homework1/mathoper/methods"
	"fmt"
)

func runOperation(name string, m *mathoper.MathOperationModule, a, b float64) {
	res, err := m.Calculate(a, b)
	if err != nil {
		fmt.Printf("Операция %s с аргументами (%.2f, %.2f), ошибка: %v\n", name, a, b, err)
		return
	}
	fmt.Printf("Операция %s с аргументами (%.2f, %.2f), результат: %.2f\n", name, a, b, res)
}

func main() {
	// Addition
	add := methods.NewAddition()
	addModule := mathoper.NewMathOperationModule(add)
	runOperation("Addition", addModule, 10, 5)   // корректно
	runOperation("Addition", addModule, 2001, 0) // некорректно — аргумент вне диапазона

	// Subtraction
	sub := methods.NewSubtraction()
	subModule := mathoper.NewMathOperationModule(sub)
	runOperation("Subtraction", subModule, 10, 5)    // корректно
	runOperation("Subtraction", subModule, -2001, 0) // некорректно — аргумент вне диапазона

	// Multiplication
	mul := methods.NewMultiplication()
	mulModule := mathoper.NewMathOperationModule(mul)
	runOperation("Multiplication", mulModule, 10, 5)   // корректно
	runOperation("Multiplication", mulModule, 1001, 2) // некорректно — аргумент вне диапазона

	// Division
	div := methods.NewDivision()
	divModule := mathoper.NewMathOperationModule(div)
	runOperation("Division", divModule, 10, 2)   // корректно
	runOperation("Division", divModule, 10, 0)   // некорректно — деление на ноль
	runOperation("Division", divModule, 2000, 1) // некорректно — аргумент вне диапазона
}
