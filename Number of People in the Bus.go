package main

import "fmt"

func Number(busStops [][2]int) int {
	var inbus int
	var outbus int
	for i := 0; i < len(busStops); i++ {
		inbus = inbus + busStops[i][0]
		outbus = outbus + busStops[i][1]
	}
	return inbus - outbus // Good Luck!
}

func main() {
	var in = [][2]int{{1, 1}, {2, 2}, {5, 0}}
	fmt.Println(Number(in))
	// var out = [][2]int{{1, 1}, {2, 2}, {5, 0}}
	// var inbus int
	// var outbus int
	// for i := 0; i < len(out); i++ {
	// 	inbus = inbus + out[i][0]
	// 	outbus = outbus + out[i][1]
	// 	fmt.Println(out[i])
	// }
	// fmt.Println("Celkem do busu: ", inbus)
	// fmt.Println("Celkem z busu: ", outbus)
}
