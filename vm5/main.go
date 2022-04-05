package main

import (
	"vm5/config"
	"vm5/slae"
	"vm5/util"
)

func main() {
	debug, _ := slae.Choleski(
		util.MakeFu64Matrix(config.MatrixDebug, 3),
		util.MakeFu64Arr(config.FreeDebug, 3),
	)
	util.PrintFu64Array(debug)

	task, _ := slae.Choleski(
		util.MakeFu64Matrix(config.Matrix, 5),
		util.MakeFu64Arr(config.Free, 5),
	)
	util.PrintFu64Array(task)

	slae.CheckSolution(
		util.MakeFu64Matrix(config.Matrix, 5),
		util.MakeFu64Arr(config.Free, 5),
		task,
	)

	/*
		То ли функция то ли калькулятор выдаёт неверное значение.
		В калькуляторе значения сумм слева: 1, 2, 3.0024..., 4.0017, 4.9696...
	*/
}
