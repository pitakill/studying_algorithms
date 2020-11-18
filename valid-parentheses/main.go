package main

import "fmt"

func main() {
	var s string

	s = "("
	fmt.Println(isValid(s))

	s = "()"
	fmt.Println(isValid(s))

	s = "()[]{}"
	fmt.Println(isValid(s))

	s = "(]"
	fmt.Println(isValid(s))

	s = "([)]"
	fmt.Println(isValid(s))

	s = "{[]}"
	fmt.Println(isValid(s))

	s = "{[]}()"
	fmt.Println(isValid(s))
}

type stack struct {
	storage []byte
}

func (s *stack) add(b byte) {
	s.storage = append(s.storage, b)
}

func (s *stack) empty() bool {
	return len(s.storage) == 0
}

func (s *stack) peek() byte {
	if s.empty() {
		return 0
	}

	return s.storage[len(s.storage)-1]
}

func (s *stack) pop() {
	if s.empty() {
		return
	}

	s.storage = s.storage[:len(s.storage)-1]
}

func isValid(s string) bool {
	// Early optimization
	if len(s)&1 == 1 {
		return false
	}

	st := new(stack)

	for i := range s {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			st.add(s[i])
			continue
		}

		if st.empty() ||
			(s[i] == ')' && st.peek() != '(') ||
			(s[i] == ']' && st.peek() != '[') ||
			(s[i] == '}' && st.peek() != '{') {
			return false
		}

		st.pop()
	}

	return st.empty()
}
