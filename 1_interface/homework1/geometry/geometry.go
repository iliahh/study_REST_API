package geometry

import "fmt"

type GeometricShape interface {
	Area() float64
	Perimeter() float64
	Name() string
}

type GeometryModule struct {
	geometryInfo   map[float64]GeometryInfo
	geometricShape GeometricShape
}

func NewGeometryModule(geometricShape GeometricShape) *GeometryModule {
	return &GeometryModule{
		geometryInfo:   make(map[float64]GeometryInfo),
		geometricShape: geometricShape,
	}
}

func PrintShapeInfo(s GeometricShape) {
	fmt.Println("Площадь:", s.Area())
	fmt.Println("Периметр:", s.Perimeter())
}
