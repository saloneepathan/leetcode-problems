package main

import (
	"fmt"
)

func kLengthApart(nums []int, k int) bool {
	prev := -1

	for i, v := range nums {
		if v == 1 {
			if prev != -1 && i-prev-1 < k {
				return false
			}
			prev = i
		}
	}
	return true
}

func main() {
	// Example 1
	nums1 := []int{1, 0, 0, 0, 1, 0, 0, 1}
	k1 := 2
	fmt.Println(kLengthApart(nums1, k1)) // true

	// Example 2
	nums2 := []int{1, 0, 0, 1, 0, 1}
	k2 := 2
	fmt.Println(kLengthApart(nums2, k2)) // false
}
