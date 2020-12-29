package main

import "fmt"

func main() {
	var nums []int

	nums = []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(n []int) int {
	l := len(n)

	for l > 0 {
		for i := 0; i < l-1; i++ {
			if n[i] == n[i+1] {
				for j := i + 1; j < l-1; j++ {
					n[j], n[j+1] = n[j+1], n[j]
				}
			}
		}
		l--
	}

	m := map[int]struct{}{}

	for i, v := range n {
		if _, ok := m[v]; ok {
			return i
		}
		m[v] = struct{}{}
	}

	return len(n)
}
