package main

import "fmt"

func main() {
	var nums []int

	nums = []int{1, 2, 3, 1, 1, 3}
	fmt.Println(numIdenticalPairs(nums))

	nums = []int{1, 1, 1, 1}
	fmt.Println(numIdenticalPairs(nums))

	nums = []int{1, 2, 3}
	fmt.Println(numIdenticalPairs(nums))
}

// func numIdenticalPairs(nums []int) int {
// var t int
//
// for i := range nums {
// for j := i + 1; j < len(nums); j++ {
// if nums[i] == nums[j] {
// t++
// }
// }
// }
//
// return t
// }

func numIdenticalPairs(nums []int) int {
	var t int
	var m = map[int]int{}

	for _, v := range nums {
		m[v]++
	}

	fmt.Println(m)

	return t
}
