package polynomial

import "vm7/fu64"

func (p *Polynomial) SumSelf(q *Polynomial) *Polynomial {
	lDiff := q.Length() - p.Length()
	minLength := q.Length()

	if lDiff > 0 {
		app := make([]fu64.Fu64, lDiff)
		p.coefs = append(p.coefs, app...) // extending coefs slice by length difference
		minLength = p.Length()
	}

	for i := 0; i < minLength; i++ { // adding respective coefs to each other
		p.coefs[i].SumSelf(&q.coefs[i])
	}

	return p
}

func (p *Polynomial) SubSelf(q *Polynomial) *Polynomial {
	qc := Copy(q)

	for _, coef := range qc.coefs {
		coef = *fu64.Neg(&coef)
	}

	return p.SumSelf(qc)
}

func (p *Polynomial) MulSelfByConst(c *fu64.Fu64) *Polynomial {
	for i := range p.coefs {
		p.coefs[i].MulSelf(c)
	}

	return p
}

// p(x - x1) = p*x - p*x1
func (p *Polynomial) MulSelfByMonomial(c *fu64.Fu64) *Polynomial {
	px1 := Copy(p)
	for i := range px1.coefs { // p * x1
		px1.coefs[i].MulSelf(c)
	}

	p.coefs = append(p.coefs, fu64.Fu64{})
	for i := len(p.coefs) - 1; i > 0; i-- { // p * x
		p.coefs[i] = p.coefs[i-1] // shifting the degree of each member by one
	}
	p.coefs[0] = fu64.Fu64{}

	return p.SubSelf(px1) // p*x - p*x1
}
