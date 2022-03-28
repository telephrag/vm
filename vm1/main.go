package main

import (
	"fmt"
	"vm1/config"
	"vm1/fu64"
	"vm1/polynomial"
)

func main() {
	var val = []float64{-2.345, 1.2220, 1.0098, 1.4789, 0.3870, 1.2340}
	var prec = []float64{config.N * 0.0001, 0, 0, 0, 0, 0.001}
	var pol []fu64.Fu64 = make([]fu64.Fu64, len(prec))
	for i := range pol {
		pol[i] = *fu64.New(val[i], prec[i])
	}

	p := polynomial.New(&pol) // if some value of pol is changed than p will change as well
	fmt.Printf("Polynomial: %s\n", p.ToString())

	x := config.W

	res := p.Compute(x)
	resRude := p.ComputeRude(x)
	fmt.Printf("Result via Gorner Scheme:    %+v\n", *res)
	fmt.Printf("Result via straight aproach: %+v\n", *resRude)

	fmt.Printf("Significant digits: %d\n", res.GetSignificantDigits())

	q := polynomial.DivByMonomialRude(p, config.C)
	fmt.Printf("Polynomial divided by (x - c): %s\n", q.ToString())

	var testVal = []float64{4, -4, 1}
	var testPrec = []float64{0, 0, 0}
	testCoeffs := make([]fu64.Fu64, 3)
	for i := range testCoeffs {
		testCoeffs[i] = *fu64.New(testVal[i], testPrec[i])
	}
	t := polynomial.New(&testCoeffs)
	fmt.Printf("Should be (x + 2): %s\n", polynomial.DivByMonomialRude(t, fu64.New(-2, 0)).ToString())
}
