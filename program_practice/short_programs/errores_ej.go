package main

import (
	"fmt"
)

// MyError -> definimos un nuevo tipo de error
type MyError struct{}

// Error MyError -> implementamos el método
func (e *MyError) Error() string {
	return "my error info"
}

// Divide -> podría devolver un error
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, &MyError{}
	}
	return a / b, nil
}

func main() {
	// Llamamos a la función divide que podría devolver un error
	result, err := divide(10, 0)

	// Verificamos si se produjo un error
	if err != nil {
		// Intentamos convertir el error a un tipo *MyError usando una afirmación de tipo
		if myErr, ok := err.(*MyError); ok {
			fmt.Println("Ocurrió un error:", myErr)
		} else {
			fmt.Println("Error desconocido")
		}
	} else {
		fmt.Println("El resultado es:", result)
	}
}
