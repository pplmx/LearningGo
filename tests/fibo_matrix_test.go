package tests

import (
	"math/big"
	"testing"
)

func BenchmarkFib(b *testing.B) {
	benchmarks := []struct {
		name   string
		fibNum *big.Int
	}{
		{name: "Fib_30", fibNum: big.NewInt(30)},
		{name: "Fib_70", fibNum: big.NewInt(70)},
		{name: "Fib_200", fibNum: big.NewInt(200)},
		{name: "Fib_1000", fibNum: big.NewInt(1000)},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fibonacci(bm.fibNum)
			}
		})
	}
}
