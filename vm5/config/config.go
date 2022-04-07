package config

import (
	"vm5/fu64"
)

const FloatPrintPrecision = 6

const N = 24.0 // k
const M = 6.0

var Eps = fu64.New(0.0001, 0)

var Matrix = [][]float64{ // ugly
	{12 + N, 2, M / 4, 1, 2},
	{4, 113 + N, 1, M / 10, M - 4},
	{1, 2, -24 - N, 3, 4},
	{1, 2 / N, 4, 33 + N, 4},
	{-1, 2, -3, 3 + M, -44 - N},
}

var Free = []float64{
	1, 2, 3, 4, 5,
}

var MatrixDebug0 = [][]float64{
	{-4, 1, 1},
	{1, -9, 3},
	{1, 2, -16},
}

var FreeDebug0 = []float64{
	2, 5, 13,
}

var MatrixDebug1 = [][]float64{
	{-4, 1, 1},
	{1, -9, 3},
	{1, 2, -16},
}

var FreeDebug1 = []float64{
	2, 5, 13,
}

// Choleski and Eidel methods
// Iteration methods with precision of Eps
