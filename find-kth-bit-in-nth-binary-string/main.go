package main

import (
	"fmt"
)

func findKthBit(n int, k int) byte {
	// Base case
	if n == 1 {
		return '0'
	}

	// Length of Sn = 2^n - 1
	length := (1 << n) - 1
	mid := (length / 2) + 1 // middle index

	if k == mid {
		return '1'
	}

	if k < mid {
		return findKthBit(n-1, k)
	}

	// If k > mid → mirror and invert
	mirrored := length - k + 1
	bit := findKthBit(n-1, mirrored)

	if bit == '0' {
		return '1'
	}
	return '0'
}

func main() {
	// Example 1
	n1, k1 := 3, 1
	fmt.Printf("Input: n = %d, k = %d\n", n1, k1)
	fmt.Printf("Output: %c\n\n", findKthBit(n1, k1))

	// Example 2
	n2, k2 := 4, 11
	fmt.Printf("Input: n = %d, k = %d\n", n2, k2)
	fmt.Printf("Output: %c\n", findKthBit(n2, k2))
}
