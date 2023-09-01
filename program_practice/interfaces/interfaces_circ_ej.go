//revisar por que falla

package main

import "fmt"

// definimos cual es el objeto que generamos
const (
	rectType   = "RECT"
	circleType = "CIRCLE"
)

func newGeometry(geoType string, values ...float64) Geometry {
	switch geoType {
	case rectType:
		return Rect{width: values[0], height: values[1]}
	case circleType:
		return Circle{radius: values[0]}
	}
	return nil
}

func main() {
	r := newGeometry(rectType, 2, 3)
	fmt.Println(r.area())
	fmt.Println(r.perim())
	c := newGeometry(circleType, 2)
	fmt.Println(c.area())
	fmt.Println(c.perim())
}
