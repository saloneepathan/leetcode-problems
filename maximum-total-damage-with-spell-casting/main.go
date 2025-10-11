package main

import (
	"fmt"
	"sort"
)

func maximumTotalDamage(power []int) int64 {
	count := make(map[int]int64)
	for _, p := range power {
		count[p]++
	}

	vals := make([]int, 0, len(count))
	for v := range count {
		vals = append(vals, v)
	}
	sort.Ints(vals)

	n := len(vals)
	gain := make([]int64, n)
	for i, v := range vals {
		gain[i] = int64(v) * count[v]
	}

	dp := make([]int64, n)
	dp[0] = gain[0]
	prev := 0

	for i := 1; i < n; i++ {
		// Move 'prev' to the last safe index
		for prev < i && vals[prev] <= vals[i]-3 {
			prev++
		}
		bestPrev := int64(0)
		if prev-1 >= 0 {
			bestPrev = dp[prev-1]
		}
		dp[i] = max(dp[i-1], bestPrev+gain[i])
	}

	return dp[n-1]
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maximumTotalDamage([]int{1, 1, 3, 4}))                     // 6
	fmt.Println(maximumTotalDamage([]int{7, 1, 6, 6}))                     // 13
	fmt.Println(maximumTotalDamage([]int{5, 9, 2, 10, 2, 7, 10, 9, 3, 8})) // 31
}
