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
	for i := range input {
		if i == 0{
			output[0] = input[0]
			continue
		}
		output[i] = input[i] + output[i-1]
	}
	return output
}
