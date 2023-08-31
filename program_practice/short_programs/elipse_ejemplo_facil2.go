package main

import "fmt"

// elipse porque aceptar un n variable de argumentos de tipo 'int'
func sumarNumeros(numeros ...int) int {
	suma := 0
	for _, num := range numeros {
		suma += num
	}
	return suma
}

func main() {
	resultado1 := sumarNumeros(5, 10, 15)
	fmt.Println("Suma 1:", resultado1)

	resultado2 := sumarNumeros(2, 4, 6, 8, 10)
	fmt.Println("Suma 2:", resultado2)
}
