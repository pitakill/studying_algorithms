package main

import (
	"fmt"
)

func main() {
	J, S := "aA", "aAAbbbb"
	fmt.Println(numJewelsInStones(J, S))

	J, S = "z", "ZZ"
	fmt.Println(numJewelsInStones(J, S))
}

func numJewelsInStones(J string, S string) int {
	c := 0

	for i := range J {
		for j := range S {
			if J[i] == S[j] {
				c++
			}
		}
	}

	return c
}

// func numJewelsInStones(J string, S string) int {
// m := make(map[rune]int)
// t := 0
//
// for _, v := range J {
// m[v] = 0
// }
//
// for _, v := range S {
// if _, ok := m[v]; ok {
// m[v]++
// }
// }
//
// for _, v := range m {
// t += v
// }
//
// return t
// }

// func numJewelsInStones(J string, S string) int {
// c := 0
//
// for _, v := range S {
// if strings.ContainsRune(J, v) {
// c++
// }
// }
//
// return c
// }
