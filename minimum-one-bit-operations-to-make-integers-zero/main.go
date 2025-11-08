package main

import (
	"fmt"
)

// minimumOneBitOperations returns the minimum number of operations
// to transform integer n into 0 using the defined bit operations.
func minimumOneBitOperations(n int) int {
	if n == 0 {
		return 0
	}

	// Find the index of the highest set bit
	k := 0
	for (1 << (k + 1)) <= n {
		k++
	}

	// Apply the recursive relation:
	// f(n) = 2^(k+1) - 1 - f(n ^ (1 << k))
	return (1 << (k + 1)) - 1 - minimumOneBitOperations(n^(1<<k))
}

func main() {
	// Example test cases
	tests := []int{0, 1, 3, 6, 9, 15, 25, 100}

	for _, n := range tests {
		fmt.Printf("n = %d -> Minimum operations = %d\n", n, minimumOneBitOperations(n))
	}
}
