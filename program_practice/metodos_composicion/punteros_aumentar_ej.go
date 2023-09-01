package main

import "fmt"

// Increase recibe un puntero de tipo entero
func Increase(v *int) {
	*v++
}
func main() {
	var v int = 19

	Increase(&v)
	fmt.Println("El valor de v ahora vale:", v)
}
