package main

import "fmt"

func maxSubarrayLength(nums []int, k int) int {
	freq := make(map[int]int)

	left := 0
	maxLen := 0

	for right := 0; right < len(nums); right++ {
		freq[nums[right]]++

		// Shrink the window while the current
		// element occurs more than k times.
		for freq[nums[right]] > k {
			freq[nums[left]]--
			left++
		}

		windowLen := right - left + 1
		if windowLen > maxLen {
			maxLen = windowLen
		}
	}

	return maxLen
}

func main() {
	// Example 1
	nums1 := []int{1, 2, 3, 1, 2, 3, 1, 2}
	k1 := 2
	fmt.Println(maxSubarrayLength(nums1, k1)) // 6

	// Example 2
	nums2 := []int{1, 2, 1, 2, 1, 2, 1, 2}
	k2 := 1
	fmt.Println(maxSubarrayLength(nums2, k2)) // 2

	// Example 3
	nums3 := []int{5, 5, 5, 5, 5, 5, 5}
	k3 := 4
	fmt.Println(maxSubarrayLength(nums3, k3)) // 4
}
