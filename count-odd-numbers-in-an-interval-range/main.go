package main

import (
	"fmt"
)

func countOdds(low int, high int) int {
	// Efficient O(1) solution
	return ((high + 1) / 2) - (low / 2)
}

func main() {
	// Example 1
	low1, high1 := 3, 7
	fmt.Println("Output:", countOdds(low1, high1)) // 3

	// Example 2
	low2, high2 := 8, 10
	fmt.Println("Output:", countOdds(low2, high2)) // 1
}
