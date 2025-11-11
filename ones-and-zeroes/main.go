package main

import (
	"fmt"
)

// findMaxForm returns the size of the largest subset of strs
// such that there are at most m zeros and n ones in the subset.
func findMaxForm(strs []string, m int, n int) int {
	// dp[i][j] = maximum number of strings that can be formed with at most i zeros and j ones
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, str := range strs {
		zeros, ones := countBits(str)
		// Traverse backwards to avoid recomputation issues
		for i := m; i >= zeros; i-- {
			for j := n; j >= ones; j-- {
				dp[i][j] = max(dp[i][j], 1+dp[i-zeros][j-ones])
			}
		}
	}

	return dp[m][n]
}

// countBits counts the number of 0s and 1s in the string
func countBits(s string) (zeros, ones int) {
	for _, ch := range s {
		if ch == '0' {
			zeros++
		} else {
			ones++
		}
	}
	return
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example 1
	strs1 := []string{"10", "0001", "111001", "1", "0"}
	m1, n1 := 5, 3
	fmt.Println("Output:", findMaxForm(strs1, m1, n1)) // Expected: 4

	// Example 2
	strs2 := []string{"10", "0", "1"}
	m2, n2 := 1, 1
	fmt.Println("Output:", findMaxForm(strs2, m2, n2)) // Expected: 2
}
