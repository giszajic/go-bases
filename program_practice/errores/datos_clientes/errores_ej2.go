package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func leerArchivo(nombreArchivo string) {
	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	}
	defer archivo.Close()

	// Lee el contenido completo del archivo
	datos, err := ioutil.ReadAll(archivo)
	if err != nil {
		panic("No se pudo leer el archivo")
	}

	// Imprime los datos leídos
	fmt.Println("Datos del archivo:")
	fmt.Println(string(datos))
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
		fmt.Println("Ejecución finalizada")
	}()

	leerArchivo("customers.txt")
}
