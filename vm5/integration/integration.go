package integration

import (
	"math"
	"vm5/fu64"
)

func gauss2(a, b *fu64.Fu64, f func(*fu64.Fu64) *fu64.Fu64) *fu64.Fu64 {
	ab := fu64.Sum(a, b).DivSelf(fu64.New(2, 0))
	ba := fu64.Sub(b, a).DivSelf(fu64.New(2*math.Sqrt(3), 0))

	fs := fu64.Sum(
		f(fu64.Sub(ab, ba)),
		f(fu64.Sum(ab, ba)),
	)

	res := fu64.Sub(b, a).DivSelf(fu64.New(2, 0))
	res.MulSelf(fs)

	return res
}

func Gauss(a, b *fu64.Fu64, f func(*fu64.Fu64) *fu64.Fu64, eps *fu64.Fu64) *fu64.Fu64 {

	var n float64 = 2
	prevRes := gauss2(a, b, f)
	thisRes := fu64.New(0, 0)
	delta := fu64.New(1, 0)

	for delta.MoreVal(eps) || delta.EqualsVal(eps) {

		thisRes.Value = 0
		thisRes.Prec = 0

		h := fu64.Sub(b, a).DivSelf(fu64.New(n, 0))
		left := fu64.Copy(a)
		right := fu64.Sum(left, h)

		for right.LessVal(b) || right.EqualsVal(b) {
			thisRes.SumSelf(gauss2(left, right, f))
			right.SumSelf(h)
			left.SumSelf(h)
		}

		delta = fu64.Abs(
			fu64.Sub(thisRes, prevRes),
		)
		prevRes = fu64.Copy(thisRes)

		n += 1
	}

	thisRes.Prec += eps.Value // check if addition instead of assignment can cause errors

	return thisRes
}
