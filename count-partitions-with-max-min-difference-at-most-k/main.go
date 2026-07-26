package main

import (
	"fmt"
)

func countPartitions(nums []int, k int) int {
	const mod = 1_000_000_007
	n := len(nums)

	dp := make([]int, n+1)
	pref := make([]int, n+1)
	dp[0] = 1
	pref[0] = 1

	minQ := []int{}
	maxQ := []int{}

	l := 0
	for r := 0; r < n; r++ {
		// maintain maxQ (monotonic decreasing)
		for len(maxQ) > 0 && nums[maxQ[len(maxQ)-1]] < nums[r] {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, r)

		// maintain minQ (monotonic increasing)
		for len(minQ) > 0 && nums[minQ[len(minQ)-1]] > nums[r] {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, r)

		// shrink window until valid
		for nums[maxQ[0]]-nums[minQ[0]] > k {
			if maxQ[0] == l {
				maxQ = maxQ[1:]
			}
			if minQ[0] == l {
				minQ = minQ[1:]
			}
			l++
		}

		// dp[r+1] = sum(dp[l]...dp[r])
		total := pref[r]
		if l > 0 {
			total = (total - pref[l-1] + mod) % mod
		}

		dp[r+1] = total
		pref[r+1] = (pref[r] + dp[r+1]) % mod
	}

	return dp[n]
}

func main() {
	fmt.Println(countPartitions([]int{9, 4, 1, 3, 7}, 4)) // expected 6
	fmt.Println(countPartitions([]int{3, 3, 4}, 0))       // expected 2
}
