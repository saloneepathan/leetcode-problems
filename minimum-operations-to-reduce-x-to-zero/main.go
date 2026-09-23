package main

import "fmt"

func minOperations(nums []int, x int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	target := total - x

	if target == 0 {
		return len(nums)
	}

	if target < 0 {
		return -1
	}

	left := 0
	windowSum := 0
	maxLen := -1

	for right := 0; right < len(nums); right++ {
		windowSum += nums[right]

		for left <= right && windowSum > target {
			windowSum -= nums[left]
			left++
		}

		if windowSum == target {
			length := right - left + 1
			if length > maxLen {
				maxLen = length
			}
		}
	}

	if maxLen == -1 {
		return -1
	}

	return len(nums) - maxLen
}

func main() {
	// Example 1
	nums1 := []int{1, 1, 4, 2, 3}
	x1 := 5
	fmt.Println(minOperations(nums1, x1)) // Output: 2

	// Example 2
	nums2 := []int{5, 6, 7, 8, 9}
	x2 := 4
	fmt.Println(minOperations(nums2, x2)) // Output: -1

	// Example 3
	nums3 := []int{3, 2, 20, 1, 1, 3}
	x3 := 10
	fmt.Println(minOperations(nums3, x3)) // Output: 5
}
