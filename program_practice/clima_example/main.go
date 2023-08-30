package main

import (
	"fmt"
)

func main() {
	var (
		temperature = 19.0
		humidity    = 0.45
		preassure   = 1023.6
	)

	fmt.Println(
		" Temperature: ", temperature,
		"\n Humidity: ", humidity,
		"\n Preassure: ", preassure,
	)
}
