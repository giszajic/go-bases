package main

import "fmt"

// Switch sin condición
func evaluarExpresion(expresion int) {
	switch {
	case expresion == 10:
		fmt.Println("La expresión es igual a 10")
	case expresion > 20:
		fmt.Println("La expresión es mayor que 20")
	default:
		fmt.Println("La expresión no cumple ninguna condición")
	}
}

// Switch para evaluar números
func evaluarNumero(numero int) {
	switch numero {
	case 1:
		fmt.Println("El número es 1")
	case 2:
		fmt.Println("El número es 2")
	default:
		fmt.Println("El número no es 1 ni 2")
	}
}

// Switch con declaración corta
func evaluarExpresion2(expresion int) {
	switch valor := expresion; valor {
	case 1:
		fmt.Println("La expresión es igual a 1")
	case 2:
		fmt.Println("La expresión es igual a 2")
	default:
		fmt.Println("La expresión no cumple ninguna condición")
	}
}

// Switch con múltiples casos
func evaluarNumero2(numero int) {
	switch numero {
	case 1, 3, 5:
		fmt.Println("El número es 1, 3 o 5")
	case 2, 4, 6:
		fmt.Println("El número es 2, 4 o 6")
	default:
		fmt.Println("El número no cumple ninguna condición")
	}
}

// Switch con fallthrough
func evaluarConFallthrough(numero int) {
	switch numero {
	case 1:
		fmt.Println("Este es el número 1")
		fallthrough
	case 2:
		fmt.Println("Este es el número 2")
	default:
		fmt.Println("Número desconocido")
	}
}

func main() {
	num1 := 10
	num2 := 20

	evaluarExpresion(num1)
	evaluarExpresion(num2)
	fmt.Println()

	evaluarExpresion2(1)
	evaluarExpresion2(3)
	fmt.Println()

	evaluarNumero(2)
	fmt.Println()

	evaluarNumero2(4)
	fmt.Println()

	evaluarConFallthrough(1)
	fmt.Println()
	evaluarConFallthrough(2)
}
