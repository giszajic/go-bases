package main

import (
	"fmt"
	"time"
)

// recibir como parámetro el canal en nuestra función
func procesar3(i int, c chan int) {
	fmt.Println(i, "-Inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "-Termina")
	c <- i
	//una vez que terminamos de procesar debemos enviarle
	//el valor al canal
}

func main() {
	c := make(chan int)
	go procesar3(1, c)
	fmt.Println("Termino el programa")
}

//revisar este ejercicio
