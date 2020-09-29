package main

import (
	"fmt"
)

func main() {
	input := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(input, target))

	input = []int{3, 2, 4}
	target = 6
	fmt.Println(twoSum(input, target))
}

func twoSum(nums []int, target int) []int {
	output := make([]int, 2)

	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			if target-v == nums[j] {
				output[0] = i
				output[1] = j
				continue
			}
		}
		continue
	}

	return output
}
