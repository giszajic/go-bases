package main

import "fmt"

func sumaSlice(nums []int) int {
	suma := 0
	for _, num := range nums {
		suma += num
	}
	return suma
}

func maximoSlice(nums []int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func acortarSlice(nums []int) []int {

	if len(nums) >= 6 {
		sliceAcortado := nums[2:5]
		//el nuevo slice hace referencia a las mismas posiciones de memoria que el original, no es una copia
		sliceAcortado[0] = -10
		return sliceAcortado
	}
	return nums
}

func main() {
	numeros := []int{3, 7, 1, 9, 4, 5}

	// sumar los elementos
	resultadoSuma := sumaSlice(numeros)
	fmt.Println("La suma de los elementos es:", resultadoSuma)

	// encontrar el valor máximo
	resultadoMaximo := maximoSlice(numeros)
	fmt.Println("El valor máximo en el slice es:", resultadoMaximo)

	//acortar slice
	sliceAcortado := acortarSlice(numeros)
	fmt.Println("Capacidad:", cap(numeros), "/n Longitud", len(numeros))
	fmt.Println("Capacidad:", cap(sliceAcortado), "/n Longitud", len(sliceAcortado))
	//no entiendo por que la capacidad da como output 4

	fmt.Println("Slice acortado: ", sliceAcortado)
	fmt.Println("Slice original (post slice): ", numeros)

}
