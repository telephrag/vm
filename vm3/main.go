package main

import (
	"fmt"
	"vm3/config"
	"vm3/integration"
)

func main() {
	g := integration.Gauss(config.A, config.B, config.F, config.Eps)

	fmt.Println(g)
}
