package slae

import (
	"errors"
	"fmt"
	"vm5/fu64"
)

func CheckSolution(m [][]*fu64.Fu64, f []*fu64.Fu64, s []*fu64.Fu64) (bool, error) {
	n := len(m)
	w := len(m[0])
	k := len(f)
	l := len(s)

	if n != w {
		return false, errors.New("matrix is not square")
	}

	if n != k {
		return false, errors.New("vector of free members has different dimension from matrix")
	}

	if n != l {
		return false, errors.New("vector of roots has different dimensions from matrix")
	}

	fmt.Println("Checking solution:")
	for i, root := range s {
		sum := fu64.New(0, 0)
		for j := range m {
			sum.SumSelf(fu64.Mul(m[i][j], root))
		}
		fmt.Println(sum, f[i])
	}

	return false, nil
}
