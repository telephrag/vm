package polynomial

import "vm7/fu64"

// P0 = a0; Pk = Pk-1*x + ak;
// e.g. P1 = P0*x + a1; P2 = P1*x + a2; etc.
func (p *Polynomial) Compute(x *fu64.Fu64) *fu64.Fu64 { // compute P(x)
	res := &p.coefs[len(p.coefs)-1]

	for i := len(p.coefs) - 1; i > 0; i-- {
		res = fu64.Mul(res, x)
		res = fu64.Sum(res, &p.coefs[i-1])
	}

	return res
}

func (p *Polynomial) ComputeRude(x *fu64.Fu64) *fu64.Fu64 { // head-on computation without using Gorner's scheme
	res := fu64.New(0, 0)

	for i := range p.coefs {
		temp := p.coefs[i]
		for j := 0; j < i; j++ {
			temp.MulSelf(x)
		}
		res.SumSelf(&temp)
	}

	return res
}

func DivByMonomialRude(p *Polynomial, c *fu64.Fu64) *Polynomial { // divide by (x - c) without remainder

	if p.Length() == 1 { // since we dont care about the remainder...
		return New(&[]fu64.Fu64{*fu64.New(0, 0)}) // if p's max degree is 0 we just return null
	}

	resCoefs := make([]fu64.Fu64, len(p.coefs)-1)
	resCoefs[len(resCoefs)-1] = p.coefs[len(p.coefs)-1] // q0 = p0

	temp := resCoefs[len(resCoefs)-1] // temp = q0
	for i := p.Length() - 2; i > 0; i-- {
		temp.MulSelf(c).SumSelf(&p.coefs[i-1]) // q0 * c + p1
		resCoefs[i-1] = temp                   // p1 = temp
	}

	return New(&resCoefs)
}
