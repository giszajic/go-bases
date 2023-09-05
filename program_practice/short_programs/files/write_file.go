package main

import (
	"fmt"
	"log"
	"os"
)

func createFile() {
	f, err := os.Create("hola2.txt")
	if err != nil {
		log.Fatal(err)
	}

	_, err2 := f.WriteString("Este texto proviene de nuestro programa")
	if err2 != nil {
		log.Fatal(err)
	}

	fmt.Print("Done")
}

func main() {
	createFile()
}
