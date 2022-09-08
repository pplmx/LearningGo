package tests

import (
	"errors"
	"math/big"
)

func Multiply(a int, b int) int {
	return a * b
}

func Divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("ZeroDivisionError")
	}
	return a / b, nil
}

func Fibonacci(n uint) uint {
	if n < 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func FibonacciWithoutIterative(n uint) uint {
	if n < 2 {
		return n
	}
	var a, b uint
	b = 1
	for n--; n > 0; n-- {
		a += b
		a, b = b, a
	}
	return b
}

func FibonacciForVeryVeryBigNumber(n uint) *big.Int {
	if n < 2 {
		return big.NewInt(int64(n))
	}
	a, b := big.NewInt(0), big.NewInt(1)
	for n--; n > 0; n-- {
		a.Add(a, b)
		a, b = b, a
	}
	return b
}
