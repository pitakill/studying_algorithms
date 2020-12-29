package main

import "fmt"

func main() {
	n, start := 5, 0
	fmt.Println(xorOperation(n, start))

	n, start = 4, 3
	fmt.Println(xorOperation(n, start))

	n, start = 1, 7
	fmt.Println(xorOperation(n, start))

	n, start = 10, 5
	fmt.Println(xorOperation(n, start))
}

// func xorOperation(n int, start int) int {
// nums := make([]int, n)
// t := 0
//
// for i := range nums {
// nums[i] = start + 2*i
// }
//
// for _, v := range nums {
// t ^= v
// }
//
// return t
// }

func xorOperation(n int, start int) int {
	var t int

	for i := 0; i < n; i++ {
		t ^= (start + 2*i)
	}

	return t
}
