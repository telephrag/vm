package polynomial

import (
	"strconv"
	"vm1/config"
	"vm1/fu64"
)

var superscripts = [10]string{
	"\u2070", "\u00b9", "\u00b2", "\u00b3", "\u2074", "\u2075", "\u2076", "\u2077", "\u2078", "\u2079",
}

const plusMinus = "\u00b1"

type Polynomial struct {
	coefs []fu64.Fu64
}

func New(c *[]fu64.Fu64) *Polynomial { // reconsider if you need a pointer as an argument
	if len(*c) == 0 {
		return &Polynomial{
			coefs: []fu64.Fu64{*fu64.New(0, 0)},
		}
	}

	return &Polynomial{
		coefs: *c,
	}
}

func (p *Polynomial) Length() int {
	return len(p.coefs)
}

func (p *Polynomial) IsEmpty() bool {
	return len(p.coefs) == 0
}

func getSuperscripts(n uint64) string { // returns member's order as unicode symbols
	var output string = ""
	for n >= 10 {
		output = superscripts[n%10] + output
		n = n / 10
	}
	output = superscripts[n%10] + output
	return output
}

func dontPrintPrecisionIfNull(n float64) string { // returns member's error if error is not null
	if n != 0 {
		return plusMinus + strconv.FormatFloat(n, 'f', config.FloatPrintPrecision, 64)
	}
	return ""
}

func (p *Polynomial) ToString() string {
	var output string = ""
	for i := range p.coefs[1:] {
		output = " " +
			strconv.FormatFloat(p.coefs[i+1].Value, 'f', config.FloatPrintPrecision, 64) + // coefficient
			dontPrintPrecisionIfNull(p.coefs[i+1].Prec) + // coefficient's error
			"x" + getSuperscripts(uint64(i+1)) + // x^n
			output
	}

	output += " " +
		strconv.FormatFloat(p.coefs[0].Value, 'f', config.FloatPrintPrecision, 64) +
		dontPrintPrecisionIfNull(p.coefs[0].Prec)

	return output
}

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
