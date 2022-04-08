package main

import (
	"vm7/cauchy_ode"
	"vm7/config"
)

func main() {
	cauchy_ode.CauchyODE(config.X0, config.Y0, config.H, config.F)

	cauchy_ode.CauchyODE(config.X0, config.Y0, config.HDebug, config.FDebug)
}
