package main

import "fmt"

func main() {
	input := []int{1, 2, 3, 4}
	fmt.Println(runningSum(input))

	input = []int{1, 1, 1, 1, 1}
	fmt.Println(runningSum(input))
}

func runningSum(input []int) []int {
	output := make([]int, len(input))

	for p := range input {
		sum := 0
		for i := 0; i <= p; i++ {
			sum = sum + input[i]
		}

		output[p] = sum
	}

	return output
}
