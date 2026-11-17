package main

import "fmt"

// minimumChanges splits s into k substrings so that the total number of
// letter changes needed to make every piece a semi-palindrome is minimal.
func minimumChanges(s string, k int) int {
	n := len(s)
	const inf = 1 << 30

	// cost[i][j] = min changes to make s[i..j] a semi-palindrome.
	cost := make([][]int, n)
	for i := range cost {
		cost[i] = make([]int, n)
		for j := range cost[i] {
			cost[i][j] = inf
		}
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ { // length 1 can never be a semi-palindrome
			l := j - i + 1
			best := inf
			for d := 1; d < l; d++ {
				if l%d != 0 {
					continue
				}
				// Each residue class (mod d) must be a palindrome:
				// compare mirrored pairs inside every class.
				m := l / d // size of each group
				changes := 0
				for r := 0; r < d; r++ {
					for a, b := 0, m-1; a < b; a, b = a+1, b-1 {
						if s[i+r+a*d] != s[i+r+b*d] {
							changes++
						}
					}
				}
				if changes < best {
					best = changes
				}
			}
			cost[i][j] = best
		}
	}

	// dp[p][i] = min changes splitting the first i chars into p pieces.
	dp := make([][]int, k+1)
	for p := range dp {
		dp[p] = make([]int, n+1)
		for i := range dp[p] {
			dp[p][i] = inf
		}
	}
	dp[0][0] = 0
	for p := 1; p <= k; p++ {
		for i := 2 * p; i <= n; i++ { // each piece needs length >= 2
			for j := 2 * (p - 1); j <= i-2; j++ {
				if dp[p-1][j] == inf {
					continue
				}
				if v := dp[p-1][j] + cost[j][i-1]; v < dp[p][i] {
					dp[p][i] = v
				}
			}
		}
	}
	return dp[k][n]
}

func main() {
	tests := []struct {
		s       string
		k, want int
	}{
		{"abcac", 2, 1},
		{"abcdef", 2, 2},
		{"aabbaa", 3, 0},
		{"ab", 1, 1},     // edge: smallest string
		{"aa", 1, 0},     // edge: already a palindrome
		{"abcabc", 1, 0}, // d=3 works
		{"abcd", 2, 2},
	}
	for _, t := range tests {
		got := minimumChanges(t.s, t.k)
		status := "PASS"
		if got != t.want {
			status = "FAIL"
		}
		fmt.Printf("input: s=%q k=%d | expected: %d | actual: %d | %s\n", t.s, t.k, t.want, got, status)
	}
}
