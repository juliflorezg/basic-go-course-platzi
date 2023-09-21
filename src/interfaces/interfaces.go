package interfaces

import "fmt"

type Rectangulo struct {
	Base   float64
	Altura float64
}
type Cuadrado struct {
	Base float64
}

type Figuras2D interface {
	area() float64
}

func (cuad Cuadrado) area() float64 {
	return (cuad.Base * cuad.Base)
}
func (rect Rectangulo) area() float64 {
	return (rect.Base * rect.Altura)
}

func CalcularArea(figura Figuras2D) {
	fmt.Println("Area:", figura.area())
}
