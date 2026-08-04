package main

import "fmt"

func findMissingElements(nums []int) []int {
	minVal, maxVal := nums[0], nums[0]

	// Find minimum and maximum
	for _, num := range nums {
		if num < minVal {
			minVal = num
		}
		if num > maxVal {
			maxVal = num
		}
	}

	// Store all numbers in a set
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}

	// Find missing numbers
	var result []int
	for i := minVal; i <= maxVal; i++ {
		if !set[i] {
			result = append(result, i)
		}
	}

	return result
}

func main() {
	// Test Case 1
	nums1 := []int{1, 4, 2, 5}
	fmt.Println("Input :", nums1)
	fmt.Println("Output:", findMissingElements(nums1))
	fmt.Println()

	// Test Case 2
	nums2 := []int{7, 8, 6, 9}
	fmt.Println("Input :", nums2)
	fmt.Println("Output:", findMissingElements(nums2))
	fmt.Println()

	// Test Case 3
	nums3 := []int{5, 1}
	fmt.Println("Input :", nums3)
	fmt.Println("Output:", findMissingElements(nums3))
}
