package main

import "fmt"

func main() {
	candies := []int{4, 2, 1, 1, 2}
	extraCandies := 1
	fmt.Println(kidsWithCandies(candies, extraCandies))

	candies = []int{12, 1, 12}
	extraCandies = 10
	fmt.Println(kidsWithCandies(candies, extraCandies))
}

func max(s []int) (index, value int) {
	for i, v := range s {
		if i == 0 || v > value {
			value = v
			index = i
		}
	}
	return
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	_, m := max(candies)
	output := make([]bool, len(candies))

	for i, v := range candies {
		output[i] = v+extraCandies >= m
	}

	return output
}
