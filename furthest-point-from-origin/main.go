package main

import (
	"fmt"
)

func furthestDistanceFromOrigin(moves string) int {
	L, R, blanks := 0, 0, 0

	for _, ch := range moves {
		if ch == 'L' {
			L++
		} else if ch == 'R' {
			R++
		} else {
			blanks++
		}
	}

	// Option 1: all blanks -> L
	leftHeavy := abs((L + blanks) - R)

	// Option 2: all blanks -> R
	rightHeavy := abs((R + blanks) - L)

	if leftHeavy > rightHeavy {
		return leftHeavy
	}
	return rightHeavy
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(furthestDistanceFromOrigin("L_RL__R")) // 3
	fmt.Println(furthestDistanceFromOrigin("_R__LL_")) // 5
	fmt.Println(furthestDistanceFromOrigin("_______")) // 7
}
