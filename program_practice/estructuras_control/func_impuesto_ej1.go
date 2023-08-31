/* Enunciado:
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento
de depositar el sueldo, para cumplir el objetivo es necesario crear una función que
devuelva el impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del
sueldo y si gana más de $150.000 se le descontará además un 10 % (27% en total).
*/

package main

import (
	"fmt"
)

func calcularImpuesto(sueldo float64) float64 {

	var impuesto float64

	if sueldo > 50000 && sueldo <= 150000 {
		impuesto = sueldo * 0.17
	} else if sueldo > 150000 {
		impuesto = sueldo * 0.27
	} else {
		impuesto = 0
	}

	return impuesto
}

func imprimirImpuestoFinal(sueldoInicial, impuesto float64) {

	fmt.Println("Sueldo inicial: ", sueldoInicial)
	fmt.Println("Impuesto: ", impuesto)
	fmt.Println("Sueldo final: ", sueldoInicial-impuesto)

}

func main() {

	sueldoInicial := 60000.00
	impuesto := calcularImpuesto(sueldoInicial)

	imprimirImpuestoFinal(sueldoInicial, impuesto)

}
