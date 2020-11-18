package main

import "fmt"

func main() {
	var nums []int
	var val int

	nums = []int{3, 2, 2, 3}
	val = 3
	fmt.Println(removeElement(nums, val))

	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	val = 2
	fmt.Println(removeElement(nums, val))
}

func removeElement(nums []int, val int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}
