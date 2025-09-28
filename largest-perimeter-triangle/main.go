package main

import (
	"fmt"
	"sort"
)

func largestPerimeter(nums []int) int {
	// Sort ascending
	sort.Ints(nums)

	// Check triples from the end (largest first)
	for i := len(nums) - 1; i >= 2; i-- {
		if nums[i-2]+nums[i-1] > nums[i] { // triangle inequality
			return nums[i] + nums[i-1] + nums[i-2]
		}
	}
	return 0
}

func main() {
	fmt.Println(largestPerimeter([]int{2, 1, 2}))     // Output: 5
	fmt.Println(largestPerimeter([]int{1, 2, 1, 10})) // Output: 0
	fmt.Println(largestPerimeter([]int{3, 6, 2, 3}))  // Output: 8
}
