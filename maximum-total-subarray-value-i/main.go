package main

import "fmt"

func maxTotalValue(nums []int, k int) int64 {
	minVal, maxVal := nums[0], nums[0]

	for _, x := range nums {
		if x < minVal {
			minVal = x
		}
		if x > maxVal {
			maxVal = x
		}
	}

	return int64(maxVal-minVal) * int64(k)
}

func main() {
	// Example 1
	nums1 := []int{1, 3, 2}
	k1 := 2
	fmt.Println(maxTotalValue(nums1, k1)) // Output: 4

	// Example 2
	nums2 := []int{4, 2, 5, 1}
	k2 := 3
	fmt.Println(maxTotalValue(nums2, k2)) // Output: 12
}
