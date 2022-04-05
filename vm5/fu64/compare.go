package fu64

func (f *Fu64) MoreVal(g *Fu64) bool {
	return f.Value > g.Value
}

func (f *Fu64) LessVal(g *Fu64) bool {
	return f.Value < g.Value
}

func (f *Fu64) EqualsVal(g *Fu64) bool {
	return f.Value == g.Value
}
