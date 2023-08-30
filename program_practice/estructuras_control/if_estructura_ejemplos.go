package main

import "fmt"

//archivo para probar distintas formas de estrucutrar IF

func suma1(num1, num2 int) {

	if num1 > 0 && num2 > 0 {
		suma := num1 + num2
		fmt.Println("El resultado de la suma es: ", suma)
	} else if num1 == 0 && num2 == 0 {
		fmt.Println("El resultado es 0")
	} else {
		fmt.Println("Al menos uno de los números es negativo.")
	}
}

func suma2() {

	//declaración corta
	var num3 int
	if num3 = 5; num3 > 0 {
		fmt.Println("El valor de num3 es mayor que 0")
	}

}

func main() {

	num1 := 10
	num2 := 20

	suma1(num1, num2)
	suma2()
}
