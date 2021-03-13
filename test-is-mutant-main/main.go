package main

import (
	"fmt"
	"strings"
)

func main() {
	input := []string{
		"TCGTATG",
		"AGCTGAG",
		"ATGATAA",
		"ATTTTGG",
		"AAAGTTG",
		"AGTACTG",
		"AGGGGTG",
	}
	fmt.Println(isMutant(input))
}

func isMutant(input []string) bool {
	sequence := strings.Join(input,"")
	mod := len(input)
	badsequence := 0

	// Verificación de secuencia por renglón
	for l := 0; l < mod; l++ {
		for i:=l*mod; i < (l+1)*mod-4; i++ {
			if sequence[i] == sequence[i+1] && sequence[i] == sequence[i+2] && sequence[i] == sequence[i+3]{
				badsequence++
			} else {
				continue
			}
		}
	}
	// Verificación de secuencia por columna
	for i:=0; i < mod*(mod-3)-1; i++ {
		for l := 0; l < mod; l++ {
			for i:=l*mod; i < (l+1)*mod-4; i++ {
				if sequence[i] == sequence[i+1] && sequence[i] == sequence[i+2] && sequence[i] == sequence[i+3]{
					badsequence++
				} else {
					continue
				}
			}
		}
	}
	// Verificación de secuencia por diaginales menos identidad
	for l := 0; l < mod-3; l++ {
		for i:=l*mod; i < (l+1)*mod-4; i++ {
			if sequence[i] == sequence[i+mod+1] && sequence[i] == sequence[i+2*(mod+1)] && sequence[i] == sequence[i+3*(mod+1)]{
				badsequence++
			} else {
				continue
			}
		}
	}
	// Verificación de secuencia por diaginales identidad
	for l := 0; l < mod-3; l++ {
		for i:=l*mod+3; i < (l+1)*mod-1; i++ {
			if sequence[i] == sequence[i+mod-1] && sequence[i] == sequence[i+2*(mod-1)] && sequence[i] == sequence[i+3*(mod-1)]{
				badsequence++
			} else {
				continue
			}
		}
	}
	return 2 <= badsequence
}
