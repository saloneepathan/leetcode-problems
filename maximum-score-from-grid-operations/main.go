package main

import (
	"fmt"
)

// Helper max function for int64
func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// maximumScore calculates the maximum score based on the given grid
func maximumScore(grid [][]int) int64 {
	n := len(grid)
	if n == 0 || len(grid[0]) == 0 {
		return 0
	}

	cols := len(grid[0])

	// If only one column, no transitions possible
	if cols == 1 {
		return 0
	}

	// DP dimensions based on columns
	dp := make([][][]int64, cols)
	for i := range dp {
		dp[i] = make([][]int64, n+1)
		for j := range dp[i] {
			dp[i][j] = make([]int64, n+1)
		}
	}

	prevMax := make([][]int64, n+1)
	prevSuffixMax := make([][]int64, n+1)

	for i := 0; i <= n; i++ {
		prevMax[i] = make([]int64, n+1)
		prevSuffixMax[i] = make([]int64, n+1)
	}

	// Prefix sums for each column
	colSum := make([][]int64, cols)
	for c := 0; c < cols; c++ {
		colSum[c] = make([]int64, n+1)
		for r := 1; r <= n; r++ {
			colSum[c][r] = colSum[c][r-1] + int64(grid[r-1][c])
		}
	}

	// DP transitions
	for i := 1; i < cols; i++ {
		for currH := 0; currH <= n; currH++ {
			for prevH := 0; prevH <= n; prevH++ {
				if currH <= prevH {
					extraScore := colSum[i][prevH] - colSum[i][currH]
					dp[i][currH][prevH] = max(
						dp[i][currH][prevH],
						prevSuffixMax[prevH][0]+extraScore,
					)
				} else {
					extraScore := colSum[i-1][currH] - colSum[i-1][prevH]
					dp[i][currH][prevH] = max(
						dp[i][currH][prevH],
						max(
							prevSuffixMax[prevH][currH],
							prevMax[prevH][currH]+extraScore,
						),
					)
				}
			}
		}

		// Update prefix and suffix maxima
		for currH := 0; currH <= n; currH++ {
			prevMax[currH][0] = dp[i][currH][0]

			for prevH := 1; prevH <= n; prevH++ {
				var penalty int64 = 0
				if prevH > currH {
					penalty = colSum[i][prevH] - colSum[i][currH]
				}
				prevMax[currH][prevH] = max(
					prevMax[currH][prevH-1],
					dp[i][currH][prevH]-penalty,
				)
			}

			prevSuffixMax[currH][n] = dp[i][currH][n]
			for prevH := n - 1; prevH >= 0; prevH-- {
				prevSuffixMax[currH][prevH] = max(
					prevSuffixMax[currH][prevH+1],
					dp[i][currH][prevH],
				)
			}
		}
	}

	// Final answer
	var ans int64 = 0
	lastCol := cols - 1
	for k := 0; k <= n; k++ {
		ans = max(ans, max(dp[lastCol][n][k], dp[lastCol][0][k]))
	}

	return ans
}

func main() {
	// Example test grid
	grid := [][]int{
		{3, 1, 2},
		{4, 6, 5},
		{7, 8, 9},
	}

	result := maximumScore(grid)

	fmt.Println("Maximum Score:", result)
}
