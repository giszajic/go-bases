package calc

func Add(num1, num2 float64) float64 {
	return num1 + num2
}

func Divide(num1, num2 float64) (result float64, err bool) {
	if num2 == 0 {
		err = true
		return
	}

	result = num1 / num2
	return
}
