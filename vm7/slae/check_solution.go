package slae

import (
	"errors"
	"fmt"
	"math/big"
	"vm7/fu64"

	"github.com/shopspring/decimal"
)

func truncate(f float64, unit float64) float64 {
	bf := big.NewFloat(0).SetPrec(1000).SetFloat64(f)
	bu := big.NewFloat(0).SetPrec(1000).SetFloat64(unit)

	bf.Quo(bf, bu)

	// Truncate:
	i := big.NewInt(0)
	bf.Int(i)
	bf.SetInt(i)

	f, _ = bf.Mul(bf, bu).Float64()
	return f
}

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

	fmt.Println("Checking solution using fu64:")
	for i := range m {
		sum := fu64.New(0, 0)
		for j := range m {
			// temp := fu64.Copy(root)
			// temp.Value = truncate(temp.Value, 0.000000001)
			sum.SumSelf(fu64.Mul(m[i][j], s[j]))
		}
		fmt.Println(sum, f[i])
	}
	fmt.Println()

	fmt.Println("Checking solution using decimal:")
	for i := range m {
		sum := decimal.NewFromFloat(0)
		for j := range m {
			temp := decimal.NewFromFloat(m[i][j].Value)
			r := decimal.NewFromFloat(s[j].Value)
			temp = temp.Mul(r)
			sum = sum.Add(temp)
		}

		fmt.Println(sum, sum.BigFloat().Prec())
	}

	fmt.Println()

	return false, nil
}
