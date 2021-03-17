package main

import (
	"fmt"
	"strings"
)

func main() {
	input := []string{
		"TCGTATGTCGTATGTCGTATG",
		"AGCTGAGIGCTGAGIGCTGAG",
		"ATTATAAATTATAAATTATAA",
		"CTCITGGCTCITGGCTCITGG",
		"AAAGTTCAAAGTTCAAAGTTC",
		"AGTACAGAGTACTGAGTACTG",
		"AGCGGCGAGCGTTGAGCGGTG",
		"TCGTATGTCGTATGTCGTATG",
		"AGCTGAGIGTTGAGIGCTGAG",
		"ATTATAAATTATAAATTATAA",
		"CTCITGGCTCITGGCTCITGG",
		"AAAGTTCAAAGTTCAAAGTTC",
		"AGTACTGAGTACTGAGTACTG",
		"AGCGGTGAGCGGTGAGCGGTG",
		"TCGTATGTCGTATGTCGTATG",
		"IGCTGAGIGCTGAGIGCTGAG",
		"ATTATAAATTAAAAATTATAA",
		"CTCITGGCTCITGGCTCITGG",
		"AAAGTTCAAAGTTCAAAGTTC",
		"AGTACTGAGTACTGAGTACTG",
		"AGCGGTGAGCGGTGAGCGGTG",
	}
	fmt.Println(isMutant(input))
}

func isMutant(input []string) bool {
	sequence := strings.Join(input, "")
	mod := len(input)
	badsequence := 0

	// Verificación de secuencia por renglón
	for l := 0; l < mod; l++ {
		for i := l * mod; i < (l+1)*mod-3; i++ {
			if sequence[i] == sequence[i+1] && sequence[i] == sequence[i+2] && sequence[i] == sequence[i+3] {
				badsequence++
				fmt.Printf("mutación de renglón: %s, en: %d,%d\n", sequence[i:i+4], l+1, (i%mod)+1)
			}
			continue
		}
	}
	// Verificación de secuencia por columna
	for i := 0; i < mod*(mod-3); i++ {
		if sequence[i] == sequence[i+mod] && sequence[i] == sequence[i+mod*2] && sequence[i] == sequence[i+mod*3] {
			badsequence++
			se := fmt.Sprintf("%s%s%s%s", string(sequence[i]), string(sequence[i+mod]), string(sequence[i+mod*2]), string(sequence[i+mod*3]))
			fmt.Printf("mutación de columna: %s, en: %d,%d \n", se, i/mod+1, (i%mod)+1)
		}
		continue
	}
	// Verificación de secuencia por diaginales menos identidad
	for l := 0; l < mod-3; l++ {
		for i := l * mod; i < (l+1)*mod-3; i++ {
			if sequence[i] == sequence[i+mod+1] && sequence[i] == sequence[i+2*(mod+1)] && sequence[i] == sequence[i+3*(mod+1)] {
				badsequence++
				se := fmt.Sprintf("%s%s%s%s", string(sequence[i]), string(sequence[i+mod+1]), string(sequence[i+(mod+1)*2]), string(sequence[i+(mod+1)*3]))
				fmt.Printf("mutación de menos identidad: %s, en: %d,%d \n", se, l+1, (i%mod)+1)
			}
			continue
		}
		// Verificación de secuencia por diaginales identidad
		for i := l*mod + 3; i < (l+1)*mod; i++ {
			if sequence[i] == sequence[i+mod-1] && sequence[i] == sequence[i+2*(mod-1)] && sequence[i] == sequence[i+3*(mod-1)] {
				badsequence++
				se := fmt.Sprintf("%s%s%s%s", string(sequence[i]), string(sequence[i+mod-1]), string(sequence[i+(mod-1)*2]), string(sequence[i+(mod-1)*3]))
				fmt.Printf("mutación de identidad: %s, en: %d,%d \n", se, l+1, (i%mod)+1)
			}
			continue
		}
	}
	return 2 <= badsequence
}
