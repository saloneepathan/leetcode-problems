package main

import (
	"fmt"
	"sort"
)

func minRemoval(nums []int, k int) int {
	n := len(nums)

	// Sort the array
	sort.Ints(nums)

	left := 0
	maxWindow := 1

	for right := 0; right < n; right++ {
		// Ensure balance condition: max <= min * k
		for nums[right] > nums[left]*k {
			left++
		}
		if right-left+1 > maxWindow {
			maxWindow = right - left + 1
		}
	}

	// Minimum removals = total - largest balanced subarray
	return n - maxWindow
}

func main() {
	// Example 1
	nums1 := []int{2, 1, 5}
	k1 := 2
	fmt.Println("Input:", nums1, "k =", k1)
	fmt.Println("Output:", minRemoval(nums1, k1))
	fmt.Println()

	// Example 2
	nums2 := []int{1, 6, 2, 9}
	k2 := 3
	fmt.Println("Input:", nums2, "k =", k2)
	fmt.Println("Output:", minRemoval(nums2, k2))
	fmt.Println()

	// Example 3
	nums3 := []int{4, 6}
	k3 := 2
	fmt.Println("Input:", nums3, "k =", k3)
	fmt.Println("Output:", minRemoval(nums3, k3))
}
