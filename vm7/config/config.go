package config

import (
	"vm7/fu64"
)

const FloatPrintPrecision = 6

const N = 24.0 // k
const M = 6.0 / 100
const L = 8.0 / 10

var F = func(x, y *fu64.Fu64) *fu64.Fu64 {
	l := fu64.NewFromF64(L)
	res := fu64.Div(l, fu64.Sum(x, l))

	m := fu64.NewFromF64(M)
	m.MulSelf(fu64.Sqrt(y))

	res.SumSelf(m)

	return res
}

var Eps = fu64.New(0.0001, 0)
var A = fu64.New(0, 0)
var B = fu64.New(1, 0)
var H = fu64.New(0.05, 0)

var X0 = fu64.New(0, 0)
var Y0 = fu64.New(1, 0)

var HDebug = fu64.New(0.2, 0)
var FDebug = func(x, y *fu64.Fu64) *fu64.Fu64 {
	return fu64.Sum(x, fu64.Sqrt(y))
}
