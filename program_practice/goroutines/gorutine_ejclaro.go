package main

import "fmt"

func multiplicarPorDos(num int, ch chan int) {
	res := num * 2
	ch <- res
}

func main() {
	n := 3

	ch := make(chan int)

	go multiplicarPorDos(n, ch)

	fmt.Println(<-ch)
}
