package main

import "fmt"

func numberOfPaths(grid [][]int, k int) int {
	const MOD = 1_000_000_007
	m, n := len(grid), len(grid[0])

	// dp[j][r] = number of ways to reach (current row, column j) with sum % k = r
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k)
	}

	for i := 0; i < m; i++ {
		newDp := make([][]int, n)
		for j := range newDp {
			newDp[j] = make([]int, k)
		}

		for j := 0; j < n; j++ {
			val := grid[i][j] % k

			for r := 0; r < k; r++ {
				// From top (previous row)
				if i > 0 && dp[j][r] > 0 {
					nr := (r + val) % k
					newDp[j][nr] = (newDp[j][nr] + dp[j][r]) % MOD
				}

				// From left (same row, previous column)
				if j > 0 && newDp[j-1][r] > 0 {
					nr := (r + val) % k
					newDp[j][nr] = (newDp[j][nr] + newDp[j-1][r]) % MOD
				}
			}

			// Start cell initialization
			if i == 0 && j == 0 {
				newDp[0][val] = 1
			}
		}

		dp = newDp
	}

	return dp[n-1][0] // sum % k == 0 at bottom-right
}

func main() {
	grid := [][]int{
		{5, 2, 4},
		{3, 0, 5},
		{0, 7, 2},
	}
	k := 3
	fmt.Println(numberOfPaths(grid, k)) // Output: 2
}
