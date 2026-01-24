package main

import (
	"fmt"
	"sort"
)

func minPairSum(nums []int) int {
	sort.Ints(nums)

	maxSum := 0
	l, r := 0, len(nums)-1

	for l < r {
		pairSum := nums[l] + nums[r]
		if pairSum > maxSum {
			maxSum = pairSum
		}
		l++
		r--
	}

	return maxSum
}

func main() {
	// Example 1
	nums1 := []int{3, 5, 2, 3}
	fmt.Println(minPairSum(nums1)) // Output: 7

	// Example 2
	nums2 := []int{3, 5, 4, 2, 4, 6}
	fmt.Println(minPairSum(nums2)) // Output: 8
}
