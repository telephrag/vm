package polynomial

import (
	"strconv"
	"vm3/config"
	"vm3/fu64"
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

func Copy(p *Polynomial) *Polynomial {
	copiedCoefs := make([]fu64.Fu64, len(p.coefs))
	copy(copiedCoefs, p.coefs)
	return New(&copiedCoefs)
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
