package main

import (
	"fmt"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	// If only one student is chosen, difference is always 0
	if k == 1 {
		return 0
	}

	// Sort the scores
	sort.Ints(nums)

	minDiff := int(1e9)

	// Sliding window of size k
	for i := 0; i+k-1 < len(nums); i++ {
		diff := nums[i+k-1] - nums[i]
		if diff < minDiff {
			minDiff = diff
		}
	}

	return minDiff
}

func main() {
	// Example 1
	nums1 := []int{90}
	k1 := 1
	fmt.Println(minimumDifference(nums1, k1)) // Output: 0

	// Example 2
	nums2 := []int{9, 4, 1, 7}
	k2 := 2
	fmt.Println(minimumDifference(nums2, k2)) // Output: 2
}
