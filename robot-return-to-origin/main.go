package main

import (
	"fmt"
)

func judgeCircle(moves string) bool {
	x, y := 0, 0

	for _, move := range moves {
		switch move {
		case 'U':
			y++
		case 'D':
			y--
		case 'R':
			x++
		case 'L':
			x--
		}
	}

	return x == 0 && y == 0
}

func main() {
	// Test cases
	fmt.Println(judgeCircle("UD"))     // true
	fmt.Println(judgeCircle("LL"))     // false
	fmt.Println(judgeCircle("RRDDLU")) // false
	fmt.Println(judgeCircle("UDLR"))   // true
}
