package main

import "fmt"

func main() {

	var sueldo float64 = 4000

	if sueldo <= 3000 {
		fmt.Println("Esta persona debe pagar impuestos")
	} else if sueldo <= 4000 {
		fmt.Printf("Debe pagar $%4.2f de su sueldo\n", (sueldo/100)*10)
	} else {
		fmt.Printf("Debe pagar $%4.2f de su sueldo\n", (sueldo/100)*15)
	}

}
