package main

import (
	"fmt"
)

func main() {
	var input []int
	var target int

	input = []int{2, 7, 11, 15}
	target = 9
	fmt.Println(twoSum(input, target))

	input = []int{3, 2, 4}
	target = 6
	fmt.Println(twoSum(input, target))
}

func twoSum(nums []int, target int) (output []int) {
	m := make(map[int]int)

	for i := range nums {
		c := target - nums[i]
		if _, ok := m[c]; ok {
			output = []int{m[c], i}
			break
		}
		m[nums[i]] = i
	}

	return
}
