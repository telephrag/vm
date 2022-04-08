package util

import (
	"fmt"
	"vm7/fu64"
)

func MakeFu64Arr(v []float64, n uint) []*fu64.Fu64 {
	res := make([]*fu64.Fu64, n)
	for i := range res {
		res[i] = fu64.NewFromF64(v[i])
	}

	return res
}

func MakeFu64Matrix(m [][]float64, n uint) [][]*fu64.Fu64 {
	res := make([][]*fu64.Fu64, n)

	for i := range res {
		res[i] = make([]*fu64.Fu64, n)
		for j := range res[i] {
			res[i][j] = fu64.NewFromF64(m[i][j])
		}
	}

	return res
}

func PrintFu64Array(v []*fu64.Fu64) {
	for _, elem := range v {
		fmt.Print(elem, " ")
	}
	fmt.Println()
}

func PrintFu64Matrix(m [][]*fu64.Fu64) {
	for _, row := range m {
		for _, elem := range row {
			fmt.Print(elem, " ")
		}
		fmt.Println()
	}
}
