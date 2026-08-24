package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func stoneGameVIII(stones []int) int {
	n := len(stones)

	// Calculate prefix sums.
	prefix := make([]int, n)
	prefix[0] = stones[0]

	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + stones[i]
	}

	// Start with taking all stones.
	dp := prefix[n-1]

	// Try every possible prefix from right to left.
	// i >= 1 because at least 2 stones must be taken.
	for i := n - 2; i >= 1; i-- {
		dp = max(dp, prefix[i]-dp)
	}

	return dp
}

func main() {
	tests := [][]int{
		{-1, 2, -3, 4, -5},
		{7, -6, 5, 10, 5, -2, -6},
		{-10, -12},
	}

	for _, stones := range tests {
		fmt.Printf("stones = %v -> answer = %d\n",
			stones,
			stoneGameVIII(stones),
		)
	}
}
