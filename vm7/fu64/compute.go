package fu64

import (
	"math"
)

func Abs(f *Fu64) *Fu64 {
	return New(math.Abs(f.Value), f.Prec)
}

func Neg(f *Fu64) *Fu64 {
	return New(f.Value*-1, f.Prec)
}

func (f *Fu64) ValueToInt() int64 {
	return int64(f.Value)
}

func Sum(a, b *Fu64) *Fu64 {
	return &Fu64{
		Value: a.Value + b.Value,
		Prec:  a.Prec + b.Prec,
	}
}

func Sub(a, b *Fu64) *Fu64 {
	return &Fu64{
		Value: a.Value - b.Value,
		Prec:  a.Prec + b.Prec,
	}
}

func Mul(a, b *Fu64) *Fu64 {
	return &Fu64{
		Value: a.Value * b.Value,
		Prec:  math.Abs(b.Value)*a.Prec + math.Abs(a.Value)*b.Prec,
	}
}

func Div(a, b *Fu64) *Fu64 {
	return &Fu64{
		Value: a.Value / b.Value,
		Prec:  (math.Abs(b.Value)*a.Prec + math.Abs(a.Value)*b.Prec) / math.Pow(b.Value, 2.0),
	}
}

func (f *Fu64) ToUIntPow(p uint64) *Fu64 {
	var i uint64
	res := New(1, 0)
	for i = 0; i < p; i++ {
		res.MulSelf(f)
	}
	return res
}

func Sqrt(f *Fu64) *Fu64 {
	div := math.Abs(1 / (2 * math.Sqrt(f.Value)))

	res := New(math.Sqrt(f.Value), div*f.Prec)

	return res
}
