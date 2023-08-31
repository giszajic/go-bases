package main

import (
	"fmt"
)

// For range
func ejemploForRange() {
	//iterando sobre un slice de números
	numeros := []int{1, 2, 3, 4, 5}

	// 1er valor: índice (que ignoramos usando _ )
	// 2do valor: elemento actual del slice
	for _, num := range numeros {
		fmt.Println(num)
	}

	/*
		// otro ejemplo: iterando sobre un slice de frutas con índice
		frutas := []string{"manzana", "banana", "pera"}

		for i, fruta := range frutas {
			fmt.Println(i, fruta)
		} */
}

// For con bucle infinito
func ejemploBucleInfinito() {

	sum := 0
	for {
		sum++ // repite para siempre
	}
	// nunca llega ya que el bucle es infinito

}

// For como bucle while
func ejemploBucleWhile() {

	i := 0
	for i < 5 {
		fmt.Println("Iteración:", i)
		i++
	}

	/*
	   sum := 1
	   for sum < 10 {
	       sum += sum
	   }
	   fmt.Println(sum)
	*/
}

// For con break
func ejemploForBreak() {

	for i := 0; i < 10; i++ {
		fmt.Println("Iteración:", i)
		if i == 5 {
			break
		}
	}

	/*
		i := 0
		for {
			fmt.Println("Iteración:", i)
			i++
			if i == 5 {
				break
			}
		} */

}

// For con continue
func ejemploForContinue() {
	for i := 0; i < 5; i++ {
		if i == 2 {
			fmt.Println("Saltando iteración:", i)
			continue
		}
		fmt.Println("Iteración:", i)
	}

	/*
		for i := 0; i < 10; i++{
		   if i % 2 == 0 {
		       continue
		   }
		   fmt.Println(i, "is odd")
		}
	*/
}

func main() {
	fmt.Println("Ejemplo de 'For range':")
	ejemploForRange()

	fmt.Println("\nEjemplo de 'For como bucle while':")
	ejemploBucleWhile()

	fmt.Println("\nEjemplo de 'For con break':")
	ejemploForBreak()

	fmt.Println("\nEjemplo de 'For con continue':")
	ejemploForContinue()

	fmt.Println("\nEjemplo de 'For con bucle infinito':")
	ejemploBucleInfinito()
}
