package main

import (
	"fmt"
	"math"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumAmount(coins [][]int) int {
	m := len(coins)
	n := len(coins[0])

	// dp[i][j][k] = max coins reaching (i,j) with k neutralizations used
	dp := make([][][]int, m)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, 3)
			for k := 0; k < 3; k++ {
				dp[i][j][k] = math.MinInt32
			}
		}
	}

	// start position
	for k := 0; k < 3; k++ {
		if coins[0][0] >= 0 {
			dp[0][0][k] = coins[0][0]
		} else {
			if k > 0 {
				dp[0][0][k] = 0
			} else {
				dp[0][0][k] = coins[0][0]
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < 3; k++ {

				if i == 0 && j == 0 {
					continue
				}

				best := math.MinInt32

				if i > 0 {
					best = max(best, dp[i-1][j][k])
				}
				if j > 0 {
					best = max(best, dp[i][j-1][k])
				}

				if best == math.MinInt32 {
					continue
				}

				// normal take
				dp[i][j][k] = max(dp[i][j][k], best+coins[i][j])

				// neutralize robber
				if coins[i][j] < 0 && k > 0 {
					bestPrev := math.MinInt32

					if i > 0 {
						bestPrev = max(bestPrev, dp[i-1][j][k-1])
					}
					if j > 0 {
						bestPrev = max(bestPrev, dp[i][j-1][k-1])
					}

					if bestPrev != math.MinInt32 {
						dp[i][j][k] = max(dp[i][j][k], bestPrev)
					}
				}
			}
		}
	}

	ans := math.MinInt32
	for k := 0; k < 3; k++ {
		ans = max(ans, dp[m-1][n-1][k])
	}

	return ans
}

func main() {

	coins1 := [][]int{
		{0, 1, -1},
		{1, -2, 3},
		{2, -3, 4},
	}

	fmt.Println(maximumAmount(coins1)) // 8

	coins2 := [][]int{
		{10, 10, 10},
		{10, 10, 10},
	}

	fmt.Println(maximumAmount(coins2)) // 40
}
