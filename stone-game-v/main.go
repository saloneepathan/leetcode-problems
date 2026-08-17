package main

import "fmt"

func stoneGameV(stoneValue []int) int {
	n := len(stoneValue)

	if n <= 1 {
		return 0
	}

	// prefix[i] = sum of stoneValue[0:i]
	prefix := make([]int, n+1)

	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + stoneValue[i]
	}

	// dp[l][r] = maximum score Alice can obtain
	// from the subarray stoneValue[l:r+1].
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	// Consider intervals from shorter to longer.
	for length := 2; length <= n; length++ {
		for l := 0; l+length <= n; l++ {
			r := l + length - 1

			for k := l; k < r; k++ {
				// Sum of left part [l...k]
				leftSum := prefix[k+1] - prefix[l]

				// Sum of right part [k+1...r]
				rightSum := prefix[r+1] - prefix[k+1]

				if leftSum < rightSum {
					// Bob discards the right side.
					// Alice keeps the left side.
					score := leftSum + dp[l][k]

					if score > dp[l][r] {
						dp[l][r] = score
					}

				} else if leftSum > rightSum {
					// Bob discards the left side.
					// Alice keeps the right side.
					score := rightSum + dp[k+1][r]

					if score > dp[l][r] {
						dp[l][r] = score
					}

				} else {
					// Equal sums: Alice can choose either side.

					leftScore := leftSum + dp[l][k]
					rightScore := rightSum + dp[k+1][r]

					if leftScore > dp[l][r] {
						dp[l][r] = leftScore
					}

					if rightScore > dp[l][r] {
						dp[l][r] = rightScore
					}
				}
			}
		}
	}

	return dp[0][n-1]
}

func main() {
	tests := []struct {
		stoneValue []int
		expected   int
	}{
		{
			stoneValue: []int{6, 2, 3, 4, 5, 5},
			expected:   18,
		},
		{
			stoneValue: []int{7, 7, 7, 7, 7, 7, 7},
			expected:   28,
		},
		{
			stoneValue: []int{4},
			expected:   0,
		},
		{
			stoneValue: []int{1, 2, 3, 4, 5},
			expected:   6,
		},
	}

	for _, test := range tests {
		result := stoneGameV(test.stoneValue)

		fmt.Printf(
			"Input: %v\nOutput: %d\nExpected: %d\n\n",
			test.stoneValue,
			result,
			test.expected,
		)
	}
}
