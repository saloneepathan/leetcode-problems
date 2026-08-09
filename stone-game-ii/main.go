func stoneGameII(piles []int) int {
	n := len(piles)

	// suffix[i] = sum of piles[i:]
	suffix := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suffix[i] = suffix[i+1] + piles[i]
	}

	// dp[i][m] = maximum stones the current player
	// can get starting at i with M = m.
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Base case:
	// If we are at or beyond the last pile, there are no stones.
	for i := n - 1; i >= 0; i-- {
		for m := n; m >= 1; m-- {

			// Can take all remaining piles.
			if i+2*m >= n {
				dp[i][m] = suffix[i]
				continue
			}

			best := 0

			// Try taking X piles, where 1 <= X <= 2*M.
			for x := 1; x <= 2*m && i+x <= n; x++ {
				nextM := m
				if x > nextM {
					nextM = x
				}

				// Current player gets everything remaining
				// except what the opponent can optimally get.
				current := suffix[i] - dp[i+x][nextM]

				if current > best {
					best = current
				}
			}

			dp[i][m] = best
		}
	}

	return dp[0][1]
}