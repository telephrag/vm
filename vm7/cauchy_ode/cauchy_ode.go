package cauchy_ode

import (
	"fmt"
	"vm7/fu64"
)

func CauchyODE(x0, y0, h *fu64.Fu64, f func(x, y *fu64.Fu64) *fu64.Fu64) { // make it return array
	k0 := fu64.New(0, 0)
	k1 := fu64.New(0, 0)
	k2 := fu64.New(0, 0)
	k3 := fu64.New(0, 0)

	res := fu64.Copy(y0) // u.i+1

	x := fu64.Copy(x0)
	y := fu64.Copy(y0) // u.i

	for x.LessVal(fu64.New(1+h.Value, 0)) {
		k0 = fu64.Mul(h, f(x, y)) // hf(xi, yi)
		k1 = fu64.Mul(h, f(       // hf(xi + h/3, yi+k0/3)
			fu64.Sum(x, fu64.Div(h, fu64.New(3, 0))),
			fu64.Sum(y, fu64.Div(k0, fu64.New(3, 0))),
		))
		k2 = fu64.Mul(h, f( // hf(xi + 2h/3, yi - k0/3 + k1)
			fu64.Sum(x, fu64.Mul(h, fu64.New(2.0/3.0, 0))),
			fu64.Sub(y, fu64.Div(k0, fu64.New(3, 0))).SumSelf(k1),
		))
		k3 = fu64.Mul(h, f( // hf(xi + h, yi + k0 - k1 + k2)
			fu64.Sum(x, h),
			fu64.Sum(y, k0).SubSelf(k1).SumSelf(k2),
		))

		ki := fu64.New(0, 0) // ki = (k0 + 3k1 + 3k2 + k3)/8
		ki.SumSelf(k0).SumSelf(k3)
		ki.SumSelf(fu64.Mul(fu64.New(3, 0), k1)).SumSelf(fu64.Mul(fu64.New(3, 0), k2))
		ki.DivSelf(fu64.New(8, 0))

		res.SumSelf(ki)
		x.SumSelf(h)
		y = fu64.Copy(res)

		fmt.Println("Y:  ", res)
	}
	fmt.Println()
}
