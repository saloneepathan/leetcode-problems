package main

import "fmt"

// numDistinct returns the number of distinct subsequences of s
// which equals t.
func numDistinct(s, t string) int {
	m, n := len(s), len(t)

	// If t is longer than s, it cannot be a subsequence.
	if m < n {
		return 0
	}

	// dp[i][j] = number of ways to form t[j:]
	// using s[i:].
	dp := make([][]int, m+1)

	for i := range dp {
		dp[i] = make([]int, n+1)

		// Empty string t[n:] can always be formed once:
		// by deleting all remaining characters from s.
		dp[i][n] = 1
	}

	// Fill the table from bottom-right to top-left.
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s[i] == t[j] {
				// Two choices:
				// 1. Use s[i] to match t[j].
				// 2. Skip s[i].
				dp[i][j] = dp[i+1][j+1] + dp[i+1][j]
			} else {
				// Characters don't match, so skip s[i].
				dp[i][j] = dp[i+1][j]
			}
		}
	}

	return dp[0][0]
}

func main() {
	tests := []struct {
		s        string
		t        string
		expected int
	}{
		{
			s:        "rabbbit",
			t:        "rabbit",
			expected: 3,
		},
		{
			s:        "babgbag",
			t:        "bag",
			expected: 5,
		},
		{
			s:        "abc",
			t:        "abc",
			expected: 1,
		},
		{
			s:        "abc",
			t:        "",
			expected: 1,
		},
		{
			s:        "",
			t:        "abc",
			expected: 0,
		},
		{
			s:        "abc",
			t:        "abcd",
			expected: 0,
		},
	}

	for _, test := range tests {
		result := numDistinct(test.s, test.t)

		fmt.Printf(
			"s = %q, t = %q\nresult = %d, expected = %d\n\n",
			test.s,
			test.t,
			result,
			test.expected,
		)
	}
}