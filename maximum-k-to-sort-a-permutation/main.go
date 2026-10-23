package main

import "fmt"

func sortPermutation(nums []int) int {
	ans := -1

	for i, num := range nums {
		if num != i {
			ans &= num
		}
	}

	if ans == -1 {
		return 0
	}

	return ans
}

func main() {
	// Test Case 1
	nums1 := []int{0, 3, 2, 1}
	fmt.Println(sortPermutation(nums1)) // Output: 1

	// Test Case 2
	nums2 := []int{0, 1, 3, 2}
	fmt.Println(sortPermutation(nums2)) // Output: 2

	// Test Case 3
	nums3 := []int{3, 2, 1, 0}
	fmt.Println(sortPermutation(nums3)) // Output: 0

	// Test Case 4
	nums4 := []int{0, 1, 2}
	fmt.Println(sortPermutation(nums4)) // Output: 0
}
