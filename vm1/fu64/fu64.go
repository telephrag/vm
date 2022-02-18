package fu64

import (
	"log"
	"math"
)

type Fu64 struct {
	Value float64
	Prec  float64
}

func New(v, p float64) *Fu64 {
	return &Fu64{
		Value: v,
		Prec:  p,
	}
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

func (f *Fu64) SumSelf(g *Fu64) *Fu64 { // return pointer to self to allow chaining e.g. f.SumSelf().MulSelf()
	f.Value += g.Value
	f.Prec += g.Prec
	return f
}

func Mul(a, b *Fu64) *Fu64 {
	return &Fu64{
		Value: a.Value * b.Value,
		Prec:  math.Abs(b.Value)*a.Prec + math.Abs(a.Value)*b.Prec,
	}
}

func (f *Fu64) MulSelf(g *Fu64) *Fu64 {
	f.Value *= g.Value
	f.Prec = math.Abs(g.Value)*f.Prec + math.Abs(f.Value)*g.Prec
	return f
}

func intPow(x int64, n int) int64 {
	var res int64 = 1
	for i := 1; i <= n; i++ {
		res *= x
	}

	return res
}

func (f *Fu64) GetSignificantDigits() int64 {
	// solve for n: prec < 0.5*10^n
	// mod := value - value % 0.5*10^n
	// compare left to right each digit of mod and value

	n := -12 // in value range of 10^12 to 10^-12 float64 is precise down to 10^-4 and up to 10^-28

	order := func(n int) float64 {
		return 0.5 * math.Pow10(n)
	}

	for order(n) < f.Prec && n != 13 {
		n++
	}
	if n == 13 {
		log.Panic("Your calculations wont be precise enough with this library. Please use a different one.\n")
	}

	mod := math.Abs(f.Value) // minus is not a significant digit

	// fmt.Println("n        = ", n)
	// fmt.Println("order(n) = ", order(n))
	// fmt.Println("mod      = ", mod)

	if n > 0 {
		return int64(mod) / intPow(10, n) // only digits before point are affected, remove n digits from right
	}
	return int64(mod * math.Pow10(n*-1)) // digits after point are affected, move point right by n and trunc the fraction
}

// NOTE: Look up why Abs() uses "bits &^ (1 << 63)" instead of "bits << 1 >> 1"
