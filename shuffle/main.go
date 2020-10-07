package main

import "fmt"

func main() {
	var input []int
	var n int

	input = []int{1, 2, 3, 4, 4, 3, 2, 1}
	n = 4
	fmt.Println(shuffle(input, n))

	input = []int{1, 1, 2, 2}
	n = 2
	fmt.Println(shuffle(input, n))
}

func shuffle(nums []int, n int) []int {
	output := make([]int, len(nums))

	for i := range nums {
		if i < n {
			output[i*2] = nums[i]
		} else {
			output[2*(i-n)+1] = nums[i]
		}
	}

	return output
}
