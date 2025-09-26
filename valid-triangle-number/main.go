package main

import (
	"fmt"
	"sort"
)

func triangleNumber(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}

	sort.Ints(nums) // Sort ascending
	count := 0

	// Fix the largest side nums[k]
	for k := n - 1; k >= 2; k-- {
		left, right := 0, k-1
		for left < right {
			if nums[left]+nums[right] > nums[k] {
				// All pairs from left..right-1 with right are valid
				count += right - left
				right--
			} else {
				left++
			}
		}
	}

	return count
}

func main() {
	fmt.Println(triangleNumber([]int{2, 2, 3, 4})) // Output: 3
	fmt.Println(triangleNumber([]int{4, 2, 3, 4})) // Output: 4
}
