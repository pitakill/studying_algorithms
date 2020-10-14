package main

import "fmt"
import "math"

func main() {
	input := int -121
	fmt.Println(isPalindrome(input))

	input = 123454321
	fmt.Println(isPalindrome(input))

	input = 98677689
	fmt.Println(isPalindrome(input))

	input = 986757689
	fmt.Println(isPalindrome(input))
}

func isPalindrome(x int) bool {
	if x < 0{
		return false
	} else {
		exp := 0
		if math.Mod(x,10)>0 {
			x = x/10 - math.Mod(x,10)
			exp = exp + 1
		}
	}
}
