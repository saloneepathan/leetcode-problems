package main

import (
	"fmt"
)

func minDeletionSize(strs []string) int {
	n := len(strs)
	m := len(strs[0])

	// dp[j] = maximum number of columns we can keep
	// such that the kept sequence ends at column j
	dp := make([]int, m)

	maxKeep := 0

	for j := 0; j < m; j++ {
		dp[j] = 1 // keep column j alone

		for i := 0; i < j; i++ {
			valid := true

			// Check if column i can come before column j
			for r := 0; r < n; r++ {
				if strs[r][i] > strs[r][j] {
					valid = false
					break
				}
			}

			if valid {
				if dp[i]+1 > dp[j] {
					dp[j] = dp[i] + 1
				}
			}
		}

		if dp[j] > maxKeep {
			maxKeep = dp[j]
		}
	}

	// Minimum deletions = total columns - max kept columns
	return m - maxKeep
}

func main() {
	tests := [][]string{
		{"babca", "bbazb"},
		{"edcba"},
		{"ghi", "def", "abc"},
	}

	for _, t := range tests {
		fmt.Println(minDeletionSize(t))
	}
}
