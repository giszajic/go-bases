package main

import "fmt"

func main() {

	// Slice inicial original
	cities := []string{"San Diego", "Los Angeles"}
	fmt.Println("Ciudades originales:", cities)

	// Slice de ciudades nuevo
	otherCities := []string{"Santa Monica", "Venice"}
	fmt.Println("Otras ciudades:", otherCities)

	// Agregamos los elementos de slice1 a slice2 con elipse
	// No se puede apendear un slice a otro, con elipse 'se desempaca'
	cities = append(cities, otherCities...)
	fmt.Println("Ciudades después de agregar otras ciudades:", cities)

}
