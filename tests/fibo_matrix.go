package tests

import "math/big"

type vector = []*big.Int
type matrix []vector

var (
	zero = new(big.Int)
	one  = big.NewInt(1)
)

func (m matrix) mul(m2 matrix) matrix {
	rows1, cols1 := len(m), len(m[0])
	rows2, cols2 := len(m2), len(m2[0])
	if cols1 != rows2 {
		panic("Matrices cannot be multiplied.")
	}
	result := make(matrix, rows1)
	temp := new(big.Int)
	for i := 0; i < rows1; i++ {
		result[i] = make(vector, cols2)
		for j := 0; j < cols2; j++ {
			result[i][j] = new(big.Int)
			for k := 0; k < rows2; k++ {
				temp.Mul(m[i][k], m2[k][j])
				result[i][j].Add(result[i][j], temp)
			}
		}
	}
	return result
}

func identityMatrix(n uint64) matrix {
	if n < 1 {
		panic("Size of identity matrix can't be less than 1")
	}
	ident := make(matrix, n)
	for i := uint64(0); i < n; i++ {
		ident[i] = make(vector, n)
		for j := uint64(0); j < n; j++ {
			ident[i][j] = new(big.Int)
			if i == j {
				ident[i][j].Set(one)
			}
		}
	}
	return ident
}

func (m matrix) pow(n *big.Int) matrix {
	le := len(m)
	if le != len(m[0]) {
		panic("Not a square matrix")
	}
	switch {
	case n.Cmp(zero) == -1:
		panic("Negative exponents not supported")
	case n.Cmp(zero) == 0:
		return identityMatrix(uint64(le))
	case n.Cmp(one) == 0:
		return m
	}
	pow := identityMatrix(uint64(le))
	base := m
	e := new(big.Int).Set(n)
	temp := new(big.Int)
	for e.Cmp(zero) > 0 {
		temp.And(e, one)
		if temp.Cmp(one) == 0 {
			pow = pow.mul(base)
		}
		e.Rsh(e, 1)
		base = base.mul(base)
	}
	return pow
}

func fibonacci(n *big.Int) *big.Int {
	if n.Cmp(zero) == 0 {
		return zero
	}
	m := matrix{{one, one}, {one, zero}}
	m = m.pow(n.Sub(n, one))
	return m[0][0]
}
