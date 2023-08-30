package main

import "fmt"

/*
func main() {
	meses := []string{"Enero", "Febrero", "Abril", "Junio", "Agosto"}

	ObtenerEstacion(meses)
}

func ObtenerEstacion(meses []string) {

	for _, mes := range meses {
		mesMinuscula := strings.ToLower(mes)
		if mesMinuscula == "enero" || mesMinuscula == "febrero" || mesMinuscula == "marzo" {
			fmt.Println("Es Verano")
		} else if mesMinuscula == "abril" || mesMinuscula == "mayo" || mesMinuscula == "junio" {
			fmt.Println("Es Otoño")
		} else if mesMinuscula == "julio" || mesMinuscula == "agosto" || mesMinuscula == "septiembre" {
			fmt.Println("Es Invierno")
		} else if mesMinuscula == "octubre" || mesMinuscula == "noviembre" || mesMinuscula == "diciembre" {
			fmt.Println("Es Primavera")
		} else {
			fmt.Println("No existe ese mes")
		}
	}
}
*/

func main() {
	var phrase string = "Hello"

	length := len(phrase)
	fmt.Println("Length: ", length, "\n")

	for i, letter := range phrase {
		fmt.Println("La letra ", i+1, " es: ", string(letter))
	}

}
