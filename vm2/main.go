package main

import (
	"fmt"
	"vm2/config"
	"vm2/fu64"
	"vm2/polynomial"
)

func main() {
	x := config.X1
	y := config.Y1

	_x := make([]fu64.Fu64, len(x))
	_y := make([]fu64.Fu64, len(y))
	for i := range _x {
		_x[i] = *fu64.New(x[i], 0)
		_y[i] = *fu64.New(y[i], 0)
	}

	p := polynomial.NewNewtonForEquipoints(&_x, &_y)

	fmt.Println(p.ToString())

}
