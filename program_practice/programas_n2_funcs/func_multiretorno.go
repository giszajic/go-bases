package main

import "fmt"

// Función que calcula la suma y el producto de dos números
func sumaYProducto(a, b int) (int, int) {
	suma := a + b
	producto := a * b
	return suma, producto
}

func main() {
	numero1 := 5
	numero2 := 7

	// Llamamos a la función sumaYProducto y recibimos los resultados en dos variables
	resultadoSuma, resultadoProducto := sumaYProducto(numero1, numero2)

	fmt.Println("Suma:\n", resultadoSuma)
	fmt.Println("Producto: \n", resultadoProducto)
}
