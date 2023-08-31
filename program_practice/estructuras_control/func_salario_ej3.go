/*
Una empresa marinera necesita calcular el salario de sus empleados basándose en la
cantidad de horas trabajadas por mes y la categoría.
Categoría C, su salario es de $1.000 por hora.
Categoría B, su salario es de $1.500 por hora, más un 20 % de su salario mensual.
Categoría A, su salario es de $3.000 por hora, más un 50 % de su salario mensual.
Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados
por mes, la categoría y que devuelva su salario.
*/

package main

import (
	"fmt"
)

func calcularSalario(categEmpleado string, cantidadHoras int, categorias map[string]int) float64 {

	salarioPorHora, existe := categorias[categEmpleado]
	if !existe {
		fmt.Println("Categoría no válida")
		return 0
	}

	salario := float64(salarioPorHora * cantidadHoras)
	switch categEmpleado {
	case "Categoría B":
		salario += float64(salarioPorHora) * 0.2
	case "Categoría A":
		salario += float64(salarioPorHora) * 0.5
	}
	return salario
}

func main() {
	categorias := map[string]int{
		"Categoría C": 1000,
		"Categoría B": 1500,
		"Categoría A": 3000,
	}
	categEmpleado := "Categoría A" // Cambiar a la categoría deseada
	cantidadHoras := 40

	salario := calcularSalario(categEmpleado, cantidadHoras, categorias)
	fmt.Println("El salario del empleado de la", categEmpleado, "es: $", salario)
}
