package structs

import "math"

//Uso de interfaces es implicito por ejemplo Rectangle tiene area que regresa un float64 ya con eso implica que puede entrar dentro de la interface de shape
type Shape interface {
	Area() float64
	Perimeter() float64
}
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radio float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radio
}
func (c Circle) Area() float64 {
	return math.Pi * c.Radio * c.Radio
}
