package main

import "fmt"

func minimumDeleteSum(s1 string, s2 string) int {
	n, m := len(s1), len(s2)

	// dp[i][j] = minimum ASCII delete sum to make s1[i:] and s2[j:] equal
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	// Base cases: when one string is exhausted
	for i := n - 1; i >= 0; i-- {
		dp[i][m] = dp[i+1][m] + int(s1[i])
	}
	for j := m - 1; j >= 0; j-- {
		dp[n][j] = dp[n][j+1] + int(s2[j])
	}

	// Fill DP table
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if s1[i] == s2[j] {
				dp[i][j] = dp[i+1][j+1]
			} else {
				deleteS1 := int(s1[i]) + dp[i+1][j]
				deleteS2 := int(s2[j]) + dp[i][j+1]
				if deleteS1 < deleteS2 {
					dp[i][j] = deleteS1
				} else {
					dp[i][j] = deleteS2
				}
			}
		}
	}

	return dp[0][0]
}

func main() {
	fmt.Println(minimumDeleteSum("sea", "eat"))     // 231
	fmt.Println(minimumDeleteSum("delete", "leet")) // 403
}
