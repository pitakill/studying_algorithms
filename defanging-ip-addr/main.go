package main

import (
	"fmt"
	"strings"
)

func main() {
	var address string
	address = "1.1.1.1"
	fmt.Println(defangIPaddr(address))

	address = "255.100.50.0"
	fmt.Println(defangIPaddr(address))
}

func defangIPaddr(address string) string {
	return strings.Join(strings.Split(address, "."), "[.]")
}
