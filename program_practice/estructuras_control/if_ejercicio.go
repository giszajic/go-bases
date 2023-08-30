package main

import (
	"fmt"
	"strings"
)

func ObtenerEstacion(meses [5]string) {

	for _, mes := range meses {
		mesMinusc := strings.ToLower(mes)
		if mesMinusc == "enero" || mesMinusc == "febrero" || mesMinusc == "marzo" {
			fmt.Println("Es Verano")
		} else if mesMinusc == "abril" || mesMinusc == "mayo" || mesMinusc == "junio" {
			fmt.Println("Es Otoño")
		} else if mesMinusc == "julio" || mesMinusc == "agosto" || mesMinusc == "septiembre" {
			fmt.Println("Es Invierno")
		} else if mesMinusc == "octubre" || mesMinusc == "noviembre" || mesMinusc == "diciembre" {
			fmt.Println("Es Primavera")
		} else {
			fmt.Println("No existe ese mes")
		}
	}
}

func main() {
	meses := [5]string{"Enero", "Febrero", "Abril", "Junio", "Agosto"}
	ObtenerEstacion(meses)
}
