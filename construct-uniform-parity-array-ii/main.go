package main

import "fmt"

func uniformArray(nums1 []int) bool {
	minVal := nums1[0]

	// Find the minimum element.
	for _, x := range nums1 {
		if x < minVal {
			minVal = x
		}
	}

	// If the minimum is odd, every other element
	// can be made odd by subtracting the minimum.
	if minVal%2 == 1 {
		return true
	}

	// If the minimum is even, we can only have a
	// uniform even array, so every element must be even.
	for _, x := range nums1 {
		if x%2 == 1 {
			return false
		}
	}

	return true
}

func main() {
	tests := [][]int{
		{1, 4, 7},
		{2, 3},
		{4, 6},
		{3, 8, 10, 15},
		{2, 4, 6, 8},
	}

	for _, nums := range tests {
		fmt.Printf("nums1 = %v -> %v\n", nums, uniformArray(nums))
	}
}