package main

import (
	"bytes"
	"fmt"
)

func main() {
	var s string

	s = "I"
	fmt.Println(romanToInt(s))

	s = "III"
	fmt.Println(romanToInt(s))

	s = "IV"
	fmt.Println(romanToInt(s))

	s = "IX"
	fmt.Println(romanToInt(s))

	s = "LVIII"
	fmt.Println(romanToInt(s))

	s = "MCMXCIV"
	fmt.Println(romanToInt(s))
}

var symbols = map[byte]struct {
	i  int
	bs []byte
}{
	'I': {1, []byte{'V', 'X'}},
	'V': {5, []byte{}},
	'X': {10, []byte{'L', 'C'}},
	'L': {50, []byte{}},
	'C': {100, []byte{'D', 'M'}},
	'D': {500, []byte{}},
	'M': {1000, []byte{}},
}

func romanToInt(s string) int {
	n := len(s)
	t := 0

	for i := 0; i < n; i++ {
		v := symbols[s[i]].bs
		j := i + 1
		if j < n && len(v) > 0 {
			r := bytes.IndexByte(v, s[j])
			if r > -1 {
				t += symbols[s[j]].i - symbols[s[i]].i
				i++
				continue
			}
		}

		t += symbols[s[i]].i
	}

	return t
}
