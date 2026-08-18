package main

import "fmt"

func largestInteger(nums []int, k int) int {
	count := make(map[int]int)

	// Count how many size-k subarrays contain each integer.
	for i := 0; i <= len(nums)-k; i++ {
		seen := make(map[int]bool)

		for j := i; j < i+k; j++ {
			seen[nums[j]] = true
		}

		for x := range seen {
			count[x]++
		}
	}

	// Find the largest integer appearing in exactly one subarray.
	ans := -1

	for x, c := range count {
		if c == 1 && x > ans {
			ans = x
		}
	}

	return ans
}

func main() {
	// Example 1
	nums := []int{3, 9, 2, 1, 7}
	k := 3

	fmt.Println(largestInteger(nums, k)) // 7

	// Example 2
	nums = []int{3, 9, 7, 2, 1, 7}
	k = 4

	fmt.Println(largestInteger(nums, k)) // 3

	// Example 3
	nums = []int{0, 0}
	k = 1

	fmt.Println(largestInteger(nums, k)) // -1
}
