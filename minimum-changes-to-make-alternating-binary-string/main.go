package main

import (
	"fmt"
)

func minOperations(s string) int {
	changes1 := 0
	changes2 := 0

	for i := 0; i < len(s); i++ {

		// pattern "010101..."
		if i%2 == 0 {
			if s[i] != '0' {
				changes1++
			}
		} else {
			if s[i] != '1' {
				changes1++
			}
		}

		// pattern "101010..."
		if i%2 == 0 {
			if s[i] != '1' {
				changes2++
			}
		} else {
			if s[i] != '0' {
				changes2++
			}
		}
	}

	if changes1 < changes2 {
		return changes1
	}
	return changes2
}

func main() {
	fmt.Println(minOperations("0100")) // 1
	fmt.Println(minOperations("10"))   // 0
	fmt.Println(minOperations("1111")) // 2
}
