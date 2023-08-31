/*
Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones. Se solicita generar una función
en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio. No se pueden introducir notas negativas.
*/

package main

import "fmt"

func calcularPromedio(notas []int) float32 {

	//evalua que no se ingresen notas negativas
	var suma int
	for _, nota := range notas {
		if nota < 0 {
			panic("No se permiten notas negativas")
		}
		suma += nota
	}

	//caso de cantidad de notas = 0
	cantidadNotas := len(notas)
	if cantidadNotas == 0 {
		return 0.0 //
	}

	promedio := float32(suma) / float32(cantidadNotas)
	return promedio

}

func main() {

	notas := []int{8, 3, 10, 7, 6, 10}

	promedio := calcularPromedio(notas)
	fmt.Println("El promedio del alumno es: ", promedio)

}
