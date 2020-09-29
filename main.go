package main

import "fmt"

func main() {
	input := []int{1, 2, 3, 4}
	fmt.Println(runningSum(input))

	input = []int{1, 1, 1, 1, 1}
	fmt.Println(runningSum(input))
}

func runningSum(input []int) []int {
	for p := range input {
		if p == 0 {
			continue
		}

		input[p] += input[p-1]
	}

	return input
}
