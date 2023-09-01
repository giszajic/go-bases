package main

import "fmt"

func main() {

	var v int = 19
	var p *int
	// Hacemos que el puntero p, referencie la dirección de memoria de la variable v
	p = &v

	fmt.Println("El puntero p referencia a la dirección de memoria: ", p)
	fmt.Println("Al desreferenciar el puntero p obtengo el valor: ", *p)

}
