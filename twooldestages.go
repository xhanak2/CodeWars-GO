package main

import "fmt"

func TwoOldestAges(ages []int) [2]int {
	var m1, m2, in int = 0, 0, 0
	for i := 0; i < len(ages); i++ {
		if m1 <= ages[i] {
			m1 = ages[i]
			in = i
		}
	}
	ages[in] = 0
	for i := 0; i < len(ages); i++ {
		if m2 <= ages[i] {
			m2 = ages[i]
		}

	}
	return [2]int{m2, m1}
}

func main() {
	var r [2]int

	r = (TwoOldestAges([]int{1, 2, 3, 4, 5, 6, 7}))
	fmt.Println(r)
}
