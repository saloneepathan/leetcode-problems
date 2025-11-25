package main

import (
	"fmt"
)

func smallestRepunitDivByK(k int) int {
	// If k has factor 2 or 5, no repunit divisible by k exists
	if k%2 == 0 || k%5 == 0 {
		return -1
	}

	r := 0
	// The length will never exceed k (Pigeonhole Principle)
	for length := 1; length <= k; length++ {
		r = (r*10 + 1) % k
		if r == 0 {
			return length
		}
	}
	return -1
}

func main() {
	// Example usage
	tests := []int{1, 2, 3, 7, 13, 17, 19}

	for _, k := range tests {
		fmt.Printf("k = %d → %d\n", k, smallestRepunitDivByK(k))
	}
}
