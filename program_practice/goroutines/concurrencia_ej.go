package main

import (
	"fmt"
	"time"
)

func main() {
	go contador("oveja")
	go contador("pez")

	time.Sleep(time.Second * 2)
}

func contador(elem string) {
	for i := 1; true; i++ {
		fmt.Println(i, elem)
		time.Sleep(time.Second / 2)
	}
}
