package fu64

import "math"

func (f *Fu64) SumSelf(g *Fu64) *Fu64 { // return pointer to self to allow chaining e.g. f.SumSelf().MulSelf()
	f.Value += g.Value
	f.Prec += g.Prec
	return f
}

func (f *Fu64) SubSelf(g *Fu64) *Fu64 { // return pointer to self to allow chaining e.g. f.SumSelf().MulSelf()
	f.Value -= g.Value
	f.Prec += g.Prec
	return f
}

func (f *Fu64) MulSelf(g *Fu64) *Fu64 {
	f.Value *= g.Value
	f.Prec = math.Abs(g.Value)*f.Prec + math.Abs(f.Value)*g.Prec
	return f
}

func (f *Fu64) DivSelf(g *Fu64) *Fu64 {
	f.Value = f.Value / g.Value
	f.Prec = (math.Abs(g.Value)*f.Prec + math.Abs(f.Value)*g.Prec) / math.Pow(g.Value, 2.0)
	return f
}
