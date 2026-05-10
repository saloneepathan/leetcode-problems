package main

import (
	"fmt"
)

func maximumJumps(nums []int, target int) int {
	n := len(nums)

	// dp[i] = maximum jumps to reach index i
	dp := make([]int, n)

	// initialize as unreachable
	for i := 0; i < n; i++ {
		dp[i] = -1
	}

	dp[0] = 0

	for i := 0; i < n; i++ {
		// skip unreachable indices
		if dp[i] == -1 {
			continue
		}

		for j := i + 1; j < n; j++ {
			diff := nums[j] - nums[i]

			if diff >= -target && diff <= target {
				if dp[i]+1 > dp[j] {
					dp[j] = dp[i] + 1
				}
			}
		}
	}

	return dp[n-1]
}

func main() {
	nums1 := []int{1, 3, 6, 4, 1, 2}
	target1 := 2
	fmt.Println(maximumJumps(nums1, target1)) // Output: 3

	nums2 := []int{1, 3, 6, 4, 1, 2}
	target2 := 3
	fmt.Println(maximumJumps(nums2, target2)) // Output: 5

	nums3 := []int{1, 3, 6, 4, 1, 2}
	target3 := 0
	fmt.Println(maximumJumps(nums3, target3)) // Output: -1
}
