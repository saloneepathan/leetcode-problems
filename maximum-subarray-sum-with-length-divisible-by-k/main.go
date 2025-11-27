package main

import (
	"fmt"
	"math"
)

func maxSubarraySum(nums []int, k int) int64 {
	n := len(nums)

	const INF int64 = math.MaxInt64
	minPref := make([]int64, k)
	for i := 0; i < k; i++ {
		minPref[i] = INF
	}

	prefix := int64(0)
	maxSum := int64(math.MinInt64)

	// prefix sum before any element (pre[0])
	minPref[0] = 0

	for i := 0; i < n; i++ {
		prefix += int64(nums[i])
		r := (i + 1) % k

		// check best possible subarray ending here
		if minPref[r] != INF {
			cand := prefix - minPref[r]
			if cand > maxSum {
				maxSum = cand
			}
		}

		// update minimum prefix for this remainder group
		if prefix < minPref[r] {
			minPref[r] = prefix
		}
	}

	return maxSum
}

func main() {
	nums := []int{-5, 1, 2, -3, 4}
	k := 2

	ans := maxSubarraySum(nums, k)
	fmt.Println(ans)
}
