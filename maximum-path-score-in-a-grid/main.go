package main

import (
	"fmt"
)

func maxPathScore(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])

	// dp[j][c] = max score reaching column j with cost c
	dp := make([][]int, n)
	for j := range dp {
		dp[j] = make([]int, k+1)
		for c := range dp[j] {
			dp[j][c] = -1
		}
	}

	// Start position
	dp[0][0] = 0

	for i := 0; i < m; i++ {
		newRow := make([][]int, n)
		for j := range newRow {
			newRow[j] = make([]int, k+1)
			for c := range newRow[j] {
				newRow[j][c] = -1
			}
		}

		for j := 0; j < n; j++ {
			// Skip start cell
			if i == 0 && j == 0 {
				newRow[j][0] = 0
				continue
			}

			cellVal := grid[i][j]
			cellScore := cellVal
			cellCost := 0
			if cellVal == 1 || cellVal == 2 {
				cellCost = 1
			}

			// From top
			if i > 0 {
				for cost := 0; cost+cellCost <= k; cost++ {
					if dp[j][cost] != -1 {
						newCost := cost + cellCost
						newScore := dp[j][cost] + cellScore
						if newScore > newRow[j][newCost] {
							newRow[j][newCost] = newScore
						}
					}
				}
			}

			// From left
			if j > 0 {
				for cost := 0; cost+cellCost <= k; cost++ {
					if newRow[j-1][cost] != -1 {
						newCost := cost + cellCost
						newScore := newRow[j-1][cost] + cellScore
						if newScore > newRow[j][newCost] {
							newRow[j][newCost] = newScore
						}
					}
				}
			}
		}

		dp = newRow
	}

	// Final answer
	ans := -1
	for cost := 0; cost <= k; cost++ {
		if dp[n-1][cost] > ans {
			ans = dp[n-1][cost]
		}
	}

	return ans
}

func main() {
	// Example 1
	grid1 := [][]int{
		{0, 1},
		{2, 0},
	}
	k1 := 1
	fmt.Println("Example 1 Output:", maxPathScore(grid1, k1)) // Expected: 2

	// Example 2
	grid2 := [][]int{
		{0, 1},
		{1, 2},
	}
	k2 := 1
	fmt.Println("Example 2 Output:", maxPathScore(grid2, k2)) // Expected: -1
}
