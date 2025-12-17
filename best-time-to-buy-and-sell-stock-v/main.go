package main

import (
	"fmt"
)

func maximumProfit(prices []int, k int) int64 {
	const negInf int64 = -1 << 60

	// dp[t][state]
	// state: 0 = flat, 1 = long, 2 = short
	dp := make([][3]int64, k+1)
	for t := 0; t <= k; t++ {
		dp[t][0] = negInf
		dp[t][1] = negInf
		dp[t][2] = negInf
	}
	dp[0][0] = 0

	for _, price := range prices {
		p := int64(price)

		next := make([][3]int64, k+1)
		for t := 0; t <= k; t++ {
			next[t][0] = dp[t][0]
			next[t][1] = dp[t][1]
			next[t][2] = dp[t][2]
		}

		for t := 0; t <= k; t++ {

			// Buy -> start long
			if dp[t][0] != negInf {
				next[t][1] = max(next[t][1], dp[t][0]-p)
			}

			// Sell -> finish long
			if t+1 <= k && dp[t][1] != negInf {
				next[t+1][0] = max(next[t+1][0], dp[t][1]+p)
			}

			// Sell -> start short
			if dp[t][0] != negInf {
				next[t][2] = max(next[t][2], dp[t][0]+p)
			}

			// Buy back -> finish short
			if t+1 <= k && dp[t][2] != negInf {
				next[t+1][0] = max(next[t+1][0], dp[t][2]-p)
			}
		}
		dp = next
	}

	var ans int64 = 0
	for t := 0; t <= k; t++ {
		ans = max(ans, dp[t][0])
	}
	return ans
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example 1
	prices1 := []int{1, 7, 9, 8, 2}
	k1 := 2
	fmt.Println("Output:", maximumProfit(prices1, k1)) // Expected: 14

	// Example 2
	prices2 := []int{12, 16, 19, 19, 8, 1, 19, 13, 9}
	k2 := 3
	fmt.Println("Output:", maximumProfit(prices2, k2)) // Expected: 36

	// Custom test
	prices3 := []int{5, 4, 3, 2, 1}
	k3 := 2
	fmt.Println("Output:", maximumProfit(prices3, k3)) // Short selling profit
}
