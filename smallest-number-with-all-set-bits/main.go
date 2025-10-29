package main

import (
	"fmt"
)

// Function to find the smallest number >= n such that its binary representation has all bits set
func smallestNumber(n int) int {
	x := 1
	for x < n {
		x = (x << 1) | 1
	}
	return x
}

func main() {
	// Example test cases
	fmt.Println(smallestNumber(5))  // Output: 7
	fmt.Println(smallestNumber(10)) // Output: 15
	fmt.Println(smallestNumber(3))  // Output: 3
	fmt.Println(smallestNumber(1))  // Output: 1
	fmt.Println(smallestNumber(16)) // Output: 31
}
