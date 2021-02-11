package lolw

import "strings"

func lengthOfLastWord(s string) int {
	sp := " "

	if s == sp {
		return 0
	}

	words := strings.Split(strings.Trim(s, sp), sp)

	return len(words[len(words)-1])
}

func lengthOfLastWord_0(s string) int {
	length := 0
	s = strings.TrimSpace(s) // testcase like "XXX ", WTF
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			break
		} else {
			length++
		}
	}
	return length
}
