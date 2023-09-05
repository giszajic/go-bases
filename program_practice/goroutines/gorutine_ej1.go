//ejecucion de forma simultanea

package main

import (
	"fmt"
	"time"
)

func procesar2(i int) {
	fmt.Println(i, "-Inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "-Termina")
}

func main() {
	for i := 0; i < 10; i++ {
		go procesar2(i)
	}
	time.Sleep(5000 * time.Millisecond)
	fmt.Println("Termino el programa")
}
