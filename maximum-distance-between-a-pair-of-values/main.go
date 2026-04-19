package main

import (
	"fmt"
)

func maxDistance(nums1 []int, nums2 []int) int {
	i, j := 0, 0
	maxDist := 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			if j-i > maxDist {
				maxDist = j - i
			}
			j++
		} else {
			i++
		}
	}

	return maxDist
}

func main() {
	// Test Case 1
	nums1 := []int{55, 30, 5, 4, 2}
	nums2 := []int{100, 20, 10, 10, 5}
	fmt.Println(maxDistance(nums1, nums2)) // Output: 2

	// Test Case 2
	nums1 = []int{2, 2, 2}
	nums2 = []int{10, 10, 1}
	fmt.Println(maxDistance(nums1, nums2)) // Output: 1

	// Test Case 3
	nums1 = []int{30, 29, 19, 5}
	nums2 = []int{25, 25, 25, 25, 25}
	fmt.Println(maxDistance(nums1, nums2)) // Output: 2
}
