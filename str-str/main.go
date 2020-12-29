package main

import "fmt"

func main() {
	var haystack, needle string

	// haystack = "hello"
	// needle = "ll"
	// fmt.Println(strStr(haystack, needle))
	//
	// haystack = "aaaaa"
	// needle = "bba"
	// fmt.Println(strStr(haystack, needle))
	//
	// haystack = ""
	// needle = ""
	// fmt.Println(strStr(haystack, needle))

	haystack = "mississippi"
	needle = "issipi"
	fmt.Println(strStr(haystack, needle))

	// haystack = "i"
	// needle = "i"
	// fmt.Println(strStr(haystack, needle))
	//
	// haystack = ""
	// needle = "i"
	// fmt.Println(strStr(haystack, needle))

	// haystack = "abc"
	// needle = "c"
	// fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	h := len(haystack)
	n := len(needle)

	if h < n {
		return -1
	}

	if h == 0 || n == 0 || haystack == needle {
		return 0
	}

	for i := range haystack {
		if haystack[i] != needle[0] {
			continue
		}

		for j := 1; j < n; j++ {
			if haystack[i+j] != needle[j] {
				fmt.Println(haystack[i : i+j])
				fmt.Println(string(haystack[i+j]), string(needle[j]))
				return -1
			}
		}

		return i
	}

	return -1
}
