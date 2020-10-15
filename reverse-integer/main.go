package main

import (
	"fmt"
	"math"
)

func main() {
	x := 123
	fmt.Println(reverse(x))

	x = -123
	fmt.Println(reverse(x))

	x = 120
	fmt.Println(reverse(x))

	x = 0
	fmt.Println(reverse(x))
}

func reverse(x int) int {
	var i int

	s := 1
	if x < 0 {
		s = -1
		x *= -1
	}

	for x > 0 {
		r := x % 10
		i = i*10 + r
		x /= 10
	}

	if i > math.MaxInt32 {
		return 0
	}

	return i * s
}
