	package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(paso1("8"))
	fmt.Println(paso1("3"))
	fmt.Println(paso1("22"))
	fmt.Println(paso1("112233"))
	fmt.Println(paso1("111223"))
	fmt.Println(paso2("111223"))
}

func paso1(s string) (int, string, string){
	cantidad := 0
	subs := s[0:]

if len(s) == 1{
	cantidad = 1
	subs = ""
} else {
	for i := 0; i < len(s); i++ {
		if s[0] == s[i] {
		cantidad = i+1
		subs = s[i+1:]
	}
}
}
return cantidad, string(s[0]), subs
}

func paso2(s string) (string){
	w, y, z := paso1(s)
	x := strconv.Itoa(w)
	cola := ""
	d, b, c := paso1(z)
	a := strconv.Itoa(d)
	if 0 < len(z){
		cola = paso2(c)
	}
	if 0 == len(z){
		cola = a + b
	}
	return x + y + cola
}

func say(n int) (string){
	if n == 1{
		return "1"
	} else {
		return paso2(say(n-1))
	}
}
