package main

import "fmt"

func resultArray(nums []int, k int) []int64 {
	result := make([]int64, k)

	// dp[r] = number of subarrays ending at the previous
	// position whose product % k == r.
	dp := make([]int64, k)

	for _, num := range nums {
		cur := make([]int64, k)

		val := num % k

		// Start a new subarray with the current element.
		cur[val]++

		// Extend all previous subarrays.
		for r := 0; r < k; r++ {
			if dp[r] > 0 {
				newRemainder := (r * val) % k
				cur[newRemainder] += dp[r]
			}
		}

		// Add all subarrays ending at the current position
		// to the final answer.
		for r := 0; r < k; r++ {
			result[r] += cur[r]
		}

		dp = cur
	}

	return result
}

func main() {
	// Example 1
	nums1 := []int{1, 2, 3, 4, 5}
	k1 := 3
	fmt.Println(resultArray(nums1, k1))
	// Output: [9 2 4]

	// Example 2
	nums2 := []int{1, 2, 4, 8, 16, 32}
	k2 := 4
	fmt.Println(resultArray(nums2, k2))
	// Output: [18 1 2 0]

	// Example 3
	nums3 := []int{1, 1, 2, 1, 1}
	k3 := 2
	fmt.Println(resultArray(nums3, k3))
	// Output: [9 6]
}