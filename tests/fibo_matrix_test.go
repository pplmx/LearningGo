package tests

import (
	"math/big"
	"testing"
)

func BenchmarkFib(b *testing.B) {
	n := big.NewInt(200)
	for i := 0; i < b.N; i++ {
		fibonacci(n)
	}
}
