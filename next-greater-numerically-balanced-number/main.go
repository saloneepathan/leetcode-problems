package main

import (
	"fmt"
)

// isBalance checks if a number is numerically balanced.
func isBalance(x int) bool {
	count := make([]int, 10)
	for x > 0 {
		count[x%10]++
		x /= 10
	}
	for i := 0; i < 10; i++ {
		if count[i] > 0 && count[i] != i {
			return false
		}
	}
	return true
}

// nextBeautifulNumber finds the next numerically balanced number after n.
func nextBeautifulNumber(n int) int {
	for i := n + 1; i <= 1224444; i++ {
		if isBalance(i) {
			return i
		}
	}
	return -1
}

func main() {
	// Test cases
	tests := []int{1, 9, 22, 3000, 3133, 34505, 1224444}

	for _, n := range tests {
		fmt.Printf("Next beautiful number after %d is %d\n", n, nextBeautifulNumber(n))
	}
}
