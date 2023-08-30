package main

import (
	"fmt"
)

func leerFrasePorLetra(frase string) {

	fmt.Println("\n Las letras de la frase son: ")

	for i, letter := range frase {
		fmt.Println(i+1, "- ", string(letter))
	}
}

func leerLargoDeFrase(frase string) {
	length := len(frase)
	fmt.Println(" Length: ", length)
}

func main() {
	var frase string = "Hello"

	leerLargoDeFrase(frase)
	leerFrasePorLetra(frase)
}
