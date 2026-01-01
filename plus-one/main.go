package main

import "fmt"

func plusOne(digits []int) []int {
	n := len(digits)

	// Start from the last digit
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		// If digit is 9, set to 0 and carry over
		digits[i] = 0
	}

	// If all digits were 9, we need one extra digit
	result := make([]int, n+1)
	result[0] = 1
	return result
}

func main() {
	// Test cases
	tests := [][]int{
		{1, 2, 3},
		{4, 3, 2, 1},
		{9},
		{9, 9, 9},
	}

	for _, test := range tests {
		fmt.Printf("Input: %v -> Output: %v\n", test, plusOne(test))
	}
}
