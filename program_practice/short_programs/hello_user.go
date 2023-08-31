package main

import "fmt"

func main() {

	//ingreso hardcodeado
	var username string
	username = "Gise Zajic"
	fmt.Println("Hello, ", username, "!")

	//ingreso por el usuario
	var username2 string
	fmt.Print("\n Your username: ")
	fmt.Scanln(&username2)
	fmt.Println("Hello, ", username2)
}
