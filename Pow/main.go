package main

import (
	"fmt"
)

func main() {
	var x float64
	var n int

	x = 2.00000
	n = 10
	fmt.Println(Pow(x, n))

	x = 2.10000
	n = 3
	fmt.Println(Pow(x, n))

	x = 2.00000
	n = -2
	fmt.Println(Pow(x, n))
}

func Pow(x float64, n int) float64 {
	var E int
	var y float64
	E = 1

	if n == 0 {
		y = 1
	}
	if n > 0{
		x = x*Pow(x,n-1)
		y = x
		}
		if n < 0{
			E = -1
			n = n*E
			y = 1 / Pow(x,n)
	}
	return y
}
