package main

import (
	"fmt"
	"math/big"
)

func hasValidPath(grid [][]byte) bool {
	n := len(grid)
	m := len(grid[0])
	pathLen := n + m - 1

	// A valid parentheses string must have even length.
	if pathLen%2 == 1 {
		return false
	}

	// Path must start with '(' and end with ')'.
	if grid[0][0] != '(' || grid[n-1][m-1] != ')' {
		return false
	}

	// dp[i][j] is a bitset.
	// Bit k represents whether balance k is possible
	// after reaching (i, j).
	dp := make([][]*big.Int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]*big.Int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = new(big.Int)
		}
	}

	// Starting balance is 1 because grid[0][0] == '('.
	dp[0][0].SetInt64(1 << 1)

	tmp := new(big.Int)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			// Skip the starting cell because it is already initialized.
			if i == 0 && j == 0 {
				continue
			}

			change := -1
			if grid[i][j] == '(' {
				change = 1
			}

			// Come from above.
			if i > 0 {
				if change == 1 {
					tmp.Lsh(dp[i-1][j], 1)
				} else {
					tmp.Rsh(dp[i-1][j], 1)
				}

				dp[i][j].Or(dp[i][j], tmp)
			}

			// Come from the left.
			if j > 0 {
				if change == 1 {
					tmp.Lsh(dp[i][j-1], 1)
				} else {
					tmp.Rsh(dp[i][j-1], 1)
				}

				dp[i][j].Or(dp[i][j], tmp)
			}
		}
	}

	// Balance 0 must be reachable at the bottom-right.
	return dp[n-1][m-1].Bit(0) == 1
}

func main() {
	grid1 := [][]byte{
		{'(', '(', '('},
		{')', '(', ')'},
		{'(', '(', ')'},
		{'(', '(', ')'},
	}

	fmt.Println(hasValidPath(grid1))

	grid2 := [][]byte{
		{'(', ')'},
		{'(', ')'},
	}

	fmt.Println(hasValidPath(grid2))
}
