package slae

import (
	"errors"
	"vm5/fu64"
	"vm5/util"
)

/*
	Solve slae via Choleski method
	matrix -- coefficients of the system
	free   -- free members vector
*/
func Choleski(matrix [][]*fu64.Fu64, free []*fu64.Fu64) ([]*fu64.Fu64, error) {
	n := len(matrix)
	w := len(matrix[0])
	m := len(free)

	if n != w {
		return nil, errors.New("matrix is not square")
	}

	if n != m {
		return nil, errors.New("vector of free member has different dimension from matrix")
	}

	temp := make([][]float64, n)
	for i := range temp {
		temp[i] = make([]float64, n)
	}
	p := util.MakeFu64Matrix(temp, uint(n))
	c := util.MakeFu64Matrix(temp, uint(n))

	for i := 0; i < n; i++ { // p.i0, c.0j
		p[i][0] = matrix[i][0]

		for j := 0; j < n; j++ {
			c[0][j] = fu64.Div(matrix[0][j], p[0][0])
		}
	}

	for i := 1; i < n; i++ { // finding P and C, matrix = PC
		for k := i; k < n; k++ {
			p[k][i].SumSelf(matrix[k][i])

			sum := fu64.New(0, 0)
			for l := 0; l < i; l++ {
				sum.SumSelf(fu64.Mul(p[i][l], c[l][i]))
			}

			p[k][i].SubSelf(sum)
		}

		for j := i; j < n; j++ {
			c[i][j].SumSelf(matrix[i][j])

			sum := fu64.New(0, 0)
			for k := 0; k < i; k++ {
				sum.SumSelf(fu64.Mul(p[i][k], c[k][j]))
			}

			c[i][j].SubSelf(sum).DivSelf(p[i][i])
		}
	}

	res := util.MakeFu64Arr(
		make([]float64, n),
		uint(n),
	)

	res[0] = fu64.Div(free[0], p[0][0]) // finding Y, PY = free
	for i := 1; i < n; i++ {
		res[i].SumSelf(free[i])

		sum := fu64.New(0, 0)
		for j := 0; j < i; j++ {
			sum.SumSelf(fu64.Mul(p[i][j], res[j]))
		}

		res[i].SubSelf(sum).DivSelf(p[i][i])
	}

	for i := n - 1; i >= 0; i-- { // findinx X, CX = Y
		sum := fu64.New(0, 0)
		for j := i + 1; j < n; j++ {
			sum.SumSelf(fu64.Mul(c[i][j], res[j]))
		}
		res[i].SubSelf(sum)
	}

	return res, nil
}
