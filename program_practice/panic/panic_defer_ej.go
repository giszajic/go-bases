package main

import (
	"fmt"
)

func main() {

	//aplicamos “defer” a la invocación de una función anónima
	defer func() {
		fmt.Println("Esta función se ejecuta a pesar de producirse panic")
	}()
	//creamos un panic con un mensaje de que se produjo
	panic("se produjo panic!!!")
}
