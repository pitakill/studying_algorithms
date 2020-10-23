package main

import (
	"fmt"
)

func main() {
	var nums []int

	nums = []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))

	nums = []int{}
	fmt.Println(threeSum(nums))

	nums = []int{0}
	fmt.Println(threeSum(nums))
}

// func threeSum(nums []int) [][]int {
// t := make([][]int, 0)
//
// final := len(nums)
//
// for i := 0; i < final; i++ {
// for j := i + 1; j < final; j++ {
// for k := j + 1; k < final; k++ {
// if nums[i]+nums[j]+nums[k] == 0 {
// a := []int{nums[i], nums[j], nums[k]}
// sort.Ints(a)
// t = append(t, a)
// }
// }
// }
// }
//
// final = len(t)
// for i := 0; i < final; i++ {
// for j := i + 1; j < final; j++ {
// if reflect.DeepEqual(t[i], t[j]) {
// t[j] = []int{}
// }
// }
// }
//
// tt := make([][]int, 0)
//
// for i := 0; i < final; i++ {
// if len(t[i]) > 0 {
// tt = append(tt, t[i])
// }
// }
//
// return tt
// }

func threeSum(nums []int) [][]int {
	m := make(map[int]int)

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			x := -(nums[i] + nums[j])
			fmt.Println(x, i, j)
			if _, ok := m[x]; ok {
				fmt.Println(x, nums[i], nums[j])
			} else {
				m[x] = nums[j]
			}
		}
	}

	fmt.Println(m)

	t := make([][]int, 0)
	return t
}
