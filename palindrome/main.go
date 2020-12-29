package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 121
	fmt.Println(isPalindrome(x))

	x = -121
	fmt.Println(isPalindrome(x))

	x = 10
	fmt.Println(isPalindrome(x))

	x = -101
	fmt.Println(isPalindrome(x))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)
	r := []rune(s)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r) == s
}
