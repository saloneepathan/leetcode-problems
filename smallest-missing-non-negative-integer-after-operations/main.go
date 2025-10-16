package main

import (
	"fmt"
)

func findSmallestInteger(nums []int, value int) int {
	counts := make(map[int]int)

	// Count how many numbers belong to each remainder modulo `value`
	for _, num := range nums {
		mod := ((num % value) + value) % value // handle negative nums correctly
		counts[mod]++
	}

	mex := 0
	for {
		mod := mex % value
		if counts[mod] > 0 {
			counts[mod]--
			mex++
		} else {
			break
		}
	}

	return mex
}

func main() {
	nums1 := []int{1, -10, 7, 13, 6, 8}
	value1 := 5
	fmt.Println(findSmallestInteger(nums1, value1)) // Output: 4

	nums2 := []int{1, -10, 7, 13, 6, 8}
	value2 := 7
	fmt.Println(findSmallestInteger(nums2, value2)) // Output: 2
}
