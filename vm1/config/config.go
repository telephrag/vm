package config

import "vm1/fu64"

const N = 24
const FloatPrintPrecision = 4

var C *fu64.Fu64 = fu64.New(0.987, N*0.0001) // task

var W *fu64.Fu64 = fu64.New(0.234*N, 0.003)  // task
var X *fu64.Fu64 = fu64.New(0.72, 0)         // 0.04...
var Y *fu64.Fu64 = fu64.New(0.234*9, 0.000)  // 77.2....
var Z *fu64.Fu64 = fu64.New(0.117*N/2, 0.02) // 1..
