package config

const FloatPrintPrecision = 6

var X1 = []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1}
var Y1 = []float64{4, 2.4, 3, 12, 5.6, 6, 16, 7.2, 10.4, 20}

var X2 = []float64{0.1, 0.2, 0.3, 0.4}
var Y2 = []float64{6, 0, 2, 6}

/*
	k = 6; m = 8; N = 20

	f1 = 1, f2 = sinx, f3 = cosx, f4 = sin2x, f5 = cos2x

	X  0.1	0.2	 0.3  0.4  0.5  0.6  0.7  0.8  0.9   1
	Y  4	2.4	 3    12   5.6  6    16   7.2  10.4  20

	1) Newton polynomial for equipoints
	2) Parabolical spline
	3) Best root-mean-square approx. function

	f1 = 1 - x, f2 = x(1 - x)^(i-1), i = 2..5, N = 3k
	f1 = 1 - x, f2 = (1 - x)x^(i-1), i = 2..5, N = 3k + 1
	f1 = 1, f2 = sinx, f3 = cosx, f4 = sin2x, f5 = cos2x

	2) 9 ranges between each pair of x we have

	{ a.k(x - x.k)^2 + b.k(x - x.k) + f(x.k), x in [x.k, x.(k+1)] }
	b.0 = 0, b.k = (2(y.k - y.(k-1)) / h.(k-1)) - b.(k-1)
	a.k = (b.(k+1) - b.k)/2h.k


	3)

*/
