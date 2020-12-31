package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(say(1))
	fmt.Println(say(2))
	fmt.Println(say(3))
	fmt.Println(say(4))
	fmt.Println(say(5))
	fmt.Println(say(6))
	fmt.Println(say(7))
	fmt.Println(say(8))
}

func paso1(s string) (string, string){
	cantidad := 0
	subs := s

if len(s) == 0 {
	return "", ""
}
if len(s) == 1 {
	cantidad = 1
	subs = ""
}
if 1 < len(s){
	for i := 1; i < len(s); i++ {
		if s[0] != s[i] {
			cantidad = i
			subs = s[i:]
			break
		} else {
		cantidad = len(s)
		subs = ""
	}
}
}
cuantos := strconv.Itoa(cantidad)+string(s[0])
return cuantos, subs
}

func paso2(s string) (string){
	y, z := paso1(s)
	cola := ""
	a, b := paso1(z)
	if 0 == len(z){
		cola = a + b
	}
	if 0 < len(z){
		cola = a + paso2(b)
	}
	return y + cola
}

func say(n int) (string){
	if n == 1 {
		return "1"
	} else {
		return paso2(say(n-1))
	}
}
