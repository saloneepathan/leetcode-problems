package main

import (
	"fmt"
)

func isOneBitCharacter(bits []int) bool {
	n := len(bits)
	i := 0

	for i < n-1 { // stop before the last bit
		if bits[i] == 1 {
			i += 2 // two-bit character
		} else {
			i++
		}
	}

	return i == n-1
}

func main() {
	// Test cases
	fmt.Println(isOneBitCharacter([]int{1, 0, 0}))    // true
	fmt.Println(isOneBitCharacter([]int{1, 1, 1, 0})) // false
	fmt.Println(isOneBitCharacter([]int{0}))          // true
	fmt.Println(isOneBitCharacter([]int{1, 0}))       // false
}
