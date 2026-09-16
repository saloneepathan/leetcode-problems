package main

import "fmt"

func numberOfSets(n int, k int) int {
	const mod = 1000000007

	dp := make([]int, n)
	prefixSums := make([]int, n+1)

	// Base case: 0 segments
	for j := 0; j < n; j++ {
		dp[j] = 1
		prefixSums[j+1] = (prefixSums[j] + dp[j]) % mod
	}

	// Add segments one by one
	for i := 1; i <= k; i++ {
		dp[0] = 0

		for j := 1; j < n; j++ {
			dp[j] = (dp[j-1] + prefixSums[j]) % mod
		}

		// Recalculate prefix sums
		for j := 0; j < n; j++ {
			prefixSums[j+1] = (prefixSums[j] + dp[j]) % mod
		}
	}

	return dp[n-1]
}

func main() {
	// Example 1
	n := 4
	k := 2
	fmt.Println(numberOfSets(n, k)) // 5

	// Example 2
	n = 3
	k = 1
	fmt.Println(numberOfSets(n, k)) // 3

	// Example 3
	n = 30
	k = 7
	fmt.Println(numberOfSets(n, k)) // 796297179
}
