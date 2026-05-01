package main

import (
	"fmt"
)

func maxRotateFunction(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	totalSum := 0
	f0 := 0

	// Compute total sum and F(0)
	for i, num := range nums {
		totalSum += num
		f0 += i * num
	}

	maxValue := f0
	current := f0

	// Compute F(k) using recurrence relation:
	// F(k) = F(k-1) + totalSum - n * nums[n-k]
	for k := 1; k < n; k++ {
		current = current + totalSum - n*nums[n-k]
		if current > maxValue {
			maxValue = current
		}
	}

	return maxValue
}

func main() {
	// Example 1
	nums1 := []int{4, 3, 2, 6}
	fmt.Println("Input:", nums1)
	fmt.Println("Maximum Rotate Function:", maxRotateFunction(nums1))

	// Example 2
	nums2 := []int{100}
	fmt.Println("Input:", nums2)
	fmt.Println("Maximum Rotate Function:", maxRotateFunction(nums2))
}
