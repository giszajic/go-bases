package main

import "fmt"

// Función  devuelve una función que suma un número constante a otro número
func crearSumador(constante int) func(int) int {
	sumador := func(numero int) int {
		return numero + constante
	}
	return sumador
}

func main() {
	// Creamos función sumadora
	sumaCinco := crearSumador(5)

	// Utilizamos la función sumadora
	resultado := sumaCinco(10)
	fmt.Println("Resultado:", resultado) // Debería imprimir 15
}
