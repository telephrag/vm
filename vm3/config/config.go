package config

import (
	"math"
	"vm3/fu64"
)

const FloatPrintPrecision = 6

const N = 24

var F func(*fu64.Fu64) *fu64.Fu64 = func(x *fu64.Fu64) *fu64.Fu64 {
	x.MulSelf(x)
	x.SumSelf(fu64.New(1, 0))
	x = fu64.Div(fu64.New(1, 0), x)
	return fu64.New(
		math.Cos(x.Value),
		math.Abs(math.Sin(x.Value)*x.Prec),
	)
}

var Eps *fu64.Fu64 = fu64.New(0.0001, 0)
var A *fu64.Fu64 = fu64.New(math.Pi/4, 0)
var B *fu64.Fu64 = fu64.New(3*math.Pi/4, 0)

// Variant 9
// f(x) = cos(1/(1+x^2)), x in [pi/4, 3pi/4], Eps = 0.0001
// Answer: 1.47795...

// Previous sums:
// 1.4793280365262502
// 1.4779922014193798
// 1.477953965005981 -- good
