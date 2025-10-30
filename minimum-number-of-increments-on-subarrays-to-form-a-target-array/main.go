package main

import (
	"fmt"
)

// minNumberOperations returns the minimum number of operations
// to form the target array from an initial zero array.
func minNumberOperations(target []int) int {
	ops := target[0]
	for i := 1; i < len(target); i++ {
		if target[i] > target[i-1] {
			ops += target[i] - target[i-1]
		}
	}
	return ops
}

func main() {
	// Sample test cases
	tests := [][]int{
		{1, 2, 3, 2, 1},
		{3, 1, 1, 2},
		{3, 1, 5, 4, 2},
		{5},             // single element
		{1, 1, 1, 1, 1}, // flat array
	}

	for _, target := range tests {
		fmt.Printf("target = %v -> min operations = %d\n", target, minNumberOperations(target))
	}
}
