package main

import "fmt"

func main() {
	var strs []string

	strs = []string{"flower", "flower"}
	fmt.Println(longestCommonPrefix(strs))

	strs = []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))

	strs = []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs))

	strs = []string{}
	fmt.Println(longestCommonPrefix(strs))

	strs = []string{"a"}
	fmt.Println(longestCommonPrefix(strs))

	strs = []string{"a", "ab"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := range strs[0] {
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) == i || strs[j][i] != strs[0][i] {
				return strs[j][:i]
			}
		}
	}

	return strs[0]
}
