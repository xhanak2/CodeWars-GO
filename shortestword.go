package main

import "fmt"

func FindShort(s string) int {
	var pocet int
	var minpocet int
	minpocet = len(s)
	for i := 0; i < len(s); i++ {
		if string(s[i]) != " " {
			pocet++
			// fmt.Println(minpocet)
		}
		if (string(s[i]) == " ") || (i == len(s)-1) {
			if minpocet > pocet {
				minpocet = pocet
				// fmt.Println(minpocet)
			}
			pocet = 0
		}
	}
	return minpocet
}

func main() {
	// var s string
	s := "Let's travel abroad shall we"

	fmt.Println(FindShort(s))
}
