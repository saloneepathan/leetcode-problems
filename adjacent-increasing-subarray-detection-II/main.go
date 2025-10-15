package main

import (
	"fmt"
)

func maxIncreasingSubarrays(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	left := make([]int, n)
	right := make([]int, n)

	// Compute lengths of increasing sequences ending at i
	left[0] = 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}

	// Compute lengths of increasing sequences starting at i
	right[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			right[i] = right[i+1] + 1
		} else {
			right[i] = 1
		}
	}

	// Find max k where both adjacent increasing subarrays exist
	res := 0
	for i := 0; i < n-1; i++ {
		k := min(left[i], right[i+1])
		if k > res {
			res = k
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Example 1
	nums1 := []int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1}
	fmt.Println("Input:", nums1)
	fmt.Println("Output:", maxIncreasingSubarrays(nums1)) // Expected: 3

	// Example 2
	nums2 := []int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7}
	fmt.Println("\nInput:", nums2)
	fmt.Println("Output:", maxIncreasingSubarrays(nums2)) // Expected: 2

	// Additional test cases
	nums3 := []int{1, 2, 3, 1, 2, 3}
	fmt.Println("\nInput:", nums3)
	fmt.Println("Output:", maxIncreasingSubarrays(nums3)) // Expected: 3

	nums4 := []int{5, 4, 3, 2, 1}
	fmt.Println("\nInput:", nums4)
	fmt.Println("Output:", maxIncreasingSubarrays(nums4)) // Expected: 0
}

