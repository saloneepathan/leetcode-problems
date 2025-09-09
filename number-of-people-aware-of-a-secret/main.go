package main

import (
	"fmt"
)

func peopleAwareOfSecret(n int, delay int, forget int) int {
	const mod = 1_000_000_007

	dp := make([]int, n+1) // dp[i] = number of people who learn on day i
	dp[1] = 1
	share := 0

	for day := 2; day <= n; day++ {
		// People who can start sharing today
		if day-delay >= 1 {
			share = (share + dp[day-delay]) % mod
		}
		// People who forget today stop sharing
		if day-forget >= 1 {
			share = (share - dp[day-forget] + mod) % mod
		}
		// Everyone currently in share group teaches new people today
		dp[day] = share
	}

	// Count those who still remember at day n
	res := 0
	for i := n - forget + 1; i <= n; i++ {
		if i > 0 {
			res = (res + dp[i]) % mod
		}
	}
	return res
}

func main() {
	fmt.Println(peopleAwareOfSecret(6, 2, 4)) // Output: 5
	fmt.Println(peopleAwareOfSecret(4, 1, 3)) // Output: 6
}
