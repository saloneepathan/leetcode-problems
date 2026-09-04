package main

import "fmt"

func firstStableIndex(nums []int, k int) int {
	n := len(nums)

	// suffixMin[i] = minimum value from nums[i] to nums[n-1]
	suffixMin := make([]int, n)
	suffixMin[n-1] = nums[n-1]

	for i := n - 2; i >= 0; i-- {
		suffixMin[i] = min(nums[i], suffixMin[i+1])
	}

	// Track maximum from nums[0] to nums[i]
	prefixMax := nums[0]

	for i := 0; i < n; i++ {
		if nums[i] > prefixMax {
			prefixMax = nums[i]
		}

		// Instability score <= k
		if prefixMax-suffixMin[i] <= k {
			return i
		}
	}

	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Example 1
	nums1 := []int{5, 0, 1, 4}
	k1 := 3
	fmt.Println(firstStableIndex(nums1, k1)) // 3

	// Example 2
	nums2 := []int{3, 2, 1}
	k2 := 1
	fmt.Println(firstStableIndex(nums2, k2)) // -1

	// Example 3
	nums3 := []int{0}
	k3 := 0
	fmt.Println(firstStableIndex(nums3, k3)) // 0
}
