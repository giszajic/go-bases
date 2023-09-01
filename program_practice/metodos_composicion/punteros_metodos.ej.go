package main

import "fmt"

type Compania struct {
	Nombre string
	Puesto string
}

type Empleado struct {
	Nombre   string
	Apellido string
	Compania Compania
}

func (e Empleado) Informacion() {
	fmt.Println(
		" Empleado:", e.Nombre, e.Apellido,
		"\n Puesto: ", e.Compania.Puesto,
		"\n Compania: ", e.Compania.Nombre,
	)
}

func (c *Compania) CambiarPuesto(nuevoPuesto string) {
	c.Puesto = nuevoPuesto
}

func main() {
	empleado := Empleado{
		Nombre:   "Juan",
		Apellido: "Lopez",
		Compania: Compania{
			Nombre: "MercadoLibre",
			Puesto: "Software Developer",
		},
	}

	empleado.Informacion()
	empleado.Compania.CambiarPuesto("Software Deveoper Ssr")
	empleado.Informacion()
}
