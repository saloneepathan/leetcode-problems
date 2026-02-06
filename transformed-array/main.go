package main

import "fmt"

func constructTransformedArray(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			result[i] = 0
			continue
		}

		// Circular index handling (works for positive & negative moves)
		newIndex := ((i+nums[i])%n + n) % n
		result[i] = nums[newIndex]
	}

	return result
}

func main() {
	// Example 1
	nums1 := []int{3, -2, 1, 1}
	fmt.Println("Input:", nums1)
	fmt.Println("Output:", constructTransformedArray(nums1))

	// Example 2
	nums2 := []int{-1, 4, -1}
	fmt.Println("Input:", nums2)
	fmt.Println("Output:", constructTransformedArray(nums2))
}
