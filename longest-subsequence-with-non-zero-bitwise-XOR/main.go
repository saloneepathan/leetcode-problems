package main

import "fmt"

func longestSubsequence(nums []int) int {
	n := len(nums)
	xor := 0
	hasNonZero := false

	for _, x := range nums {
		xor ^= x

		if x != 0 {
			hasNonZero = true
		}
	}

	// Entire array has non-zero XOR.
	if xor != 0 {
		return n
	}

	// Total XOR is zero, but at least one element is non-zero.
	// Removing that element makes the XOR non-zero.
	if hasNonZero {
		return n - 1
	}

	// All elements are zero.
	return 0
}

func main() {
	nums1 := []int{1, 2, 3}
	fmt.Println(longestSubsequence(nums1)) // 2

	nums2 := []int{2, 3, 4}
	fmt.Println(longestSubsequence(nums2)) // 3

	nums3 := []int{0, 0, 0}
	fmt.Println(longestSubsequence(nums3)) // 0
}