package tests

import "errors"

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	return a * b
}

func divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("ZeroDivisionError")
	}
	return a / b, nil
}
