package main

import (
	"Study_REST_API/1_interface/homework1/geometry"
	"Study_REST_API/1_interface/homework1/geometry/methods"
	"fmt"
)

/*
- Необходимо разработать модуль для работы с геометрическими фигурами
- Каждая фигура умеет вычислять свою площадь и периметр
- Разные фигуры реализуются по-разному, но взаимодействие с ними
 должно быть унифицировано через интерфейс
- Необходимо описать фигуры и функцию, которая сможет принимать
 различные виды фигур и высчитывать их площадь и периметр
*/

func main() {
	circle := methods.NewCircle(5)
	parallelogram := methods.NewParallelogram(4, 6, 30)
	rectangle := methods.NewRectangle(4, 6)
	rhombus := methods.NewRhombus(5, 30)
	square := methods.NewSquare(4)
	trapezoid := methods.NewTrapezoid(3, 5, 4, 6, 3)
	triangle := methods.NewTriangle(3, 4, 4, 5)

	shapes := []geometry.GeometricShape{
		circle, parallelogram,
		rectangle,
		rhombus,
		square,
		trapezoid,
		triangle,
	}

	fmt.Println("Обработка геометрических фигур:")
	for i, shape := range shapes {
		fmt.Printf("%d. %s - Площадь: %.2f, Периметр: %.2f\n",
			i+1, shape.Name(), shape.Area(), shape.Perimeter())
	}
}
