package main

import (
	"fmt"
	"log"
	"os"
)

func createFile2() {
	file := "./hola.txt"
	f, err := os.Stat(file)
	if err != err {
		log.Fatal(err)
	}

	//os.Exit(0)

	fmt.Println("Es un directorio: ", f.IsDir())
	fmt.Println("Nombre del archivo/directorio: ", f.Name())
	fmt.Println("Tamaño del archivo en bytes: ", f.Size())
	fmt.Println("Fecha y hora del archivo: ", f.ModTime())
	fmt.Println("Permisos del archivo: ", f.Mode())
}

func main() {
	createFile2()
}
