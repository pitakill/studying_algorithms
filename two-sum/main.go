package main

import "fmt"

func main() {
	input := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(input, target))

	input1 := []int{3, 2, 4}
	target1 := 6
	fmt.Println(twoSum(input1, target1))
}

func twoSum(nums []int, target int) []int {
	if len(nums) == 2 {
		return []int{0, 1}
	}
	output := make([]int, 2)
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			if target-v == nums[j] {
				output = append([]int{}, i, j)
			}
		}
	}
	return output
}
