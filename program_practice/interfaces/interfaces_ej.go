package main

import (
	"fmt"
	"math"
)

// INTERFASES
type Geometry interface {
	area() float64
	perim() float64
}

// struct Circle
type Circle struct {
	radius float64
}

// funciones que muestran area y perimetro
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// struct rectangulo
type Rect struct {
	width, height float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (r Rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// funcion que imprime area y perim
func details(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// ejecutando funcion
func main() {
	c := Circle{radius: 5}
	r := Rect{width: 3, height: 4}

	details(r)
	details(c)
}

/*
Si queremos generar mas figuras utilizando "details":
INTERFACES nos permiten implementar el mismo
comportamiento a diferentes objetos
*/
