package polynomial

import (
	"vm3/fu64"
)

func recalculateTable(t *[]fu64.Fu64) *[]fu64.Fu64 { // TODO: Add comments
	_t := *t // to make it prettier

	for i := 0; i < len(*t)-1; i++ {
		(*t)[i] = *fu64.Sub(&_t[i+1], &_t[i])
	}

	return &_t
}

// i = 0..9
// x.0 = 0.1
// y.0 = 4
// k = 0..9
//
// dy = y.(i+1) - y.i
// d^n(y) = d^n-1(y.(i+1)) - d^n-1(y.i)
//
// For i = 0..9, y.0 + (dy/h)(x-x.0) + (d^2(y)/2!h)(x-x.0)(x-x.1) + ...
// a.k = (d^k(y)/k!h^k)(x-x.0)(x-x.1)..(x-x.k-1)

func NewNewtonForEquipoints(x, y *[]fu64.Fu64) *Polynomial { // add error handling?
	_x := *x
	_y := *y

	hCurr := fu64.Sub(&_x[1], &_x[0])
	h := fu64.Copy(hCurr)

	coefs := make([]fu64.Fu64, len(_x))
	res := New(&coefs)

	for i := range _x {

		// fmt.Printf("%d)\n", i) // debug, iteration

		c := _y[0] // c is not used !!!
		// fmt.Printf("c at the beginning: %f\n", c.Value)
		for j := 1; j <= i; j++ { // j := 1 so that we wont divide by zero when dividing by k!
			c.DivSelf(h).DivSelf(fu64.New(float64(j), 0)) // divide each member by h^k * k!
		}
		// fmt.Printf("c at the end:       %f\n", c.Value)
		// fmt.Printf("hcurr:              %v\n", hCurr)

		m := New(&[]fu64.Fu64{{Value: 1, Prec: 0}})
		for k := 0; k < i; k++ {
			m.MulSelfByMonomial(fu64.Neg(&_x[k]))
		}

		// fmt.Printf("m:                  %s\n", m.ToString())

		m.MulSelfByConst(&c)

		// fmt.Printf("m * c:              %s\n", m.ToString())

		res.SumSelf(m)

		// fmt.Printf("res:                %s\n", res.ToString())

		recalculateTable(&_y) // to get new dy
		_y = _y[:len(_y)-1]

		// fmt.Println("_y: ", _y)
		// fmt.Println()
	}

	return res
}
