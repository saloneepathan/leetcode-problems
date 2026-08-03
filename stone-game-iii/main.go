package main

import (
	"fmt"
	"math"
)

func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)

	dp := make([]int, n+1)

	for i := n - 1; i >= 0; i-- {
		best := math.MinInt32
		sum := 0

		for k := 0; k < 3 && i+k < n; k++ {
			sum += stoneValue[i+k]
			if sum-dp[i+k+1] > best {
				best = sum - dp[i+k+1]
			}
		}

		dp[i] = best
	}

	if dp[0] > 0 {
		return "Alice"
	}
	if dp[0] < 0 {
		return "Bob"
	}
	return "Tie"
}

func main() {
	fmt.Println(stoneGameIII([]int{1, 2, 3, 7}))  // Bob
	fmt.Println(stoneGameIII([]int{1, 2, 3, -9})) // Alice
	fmt.Println(stoneGameIII([]int{1, 2, 3, 6}))  // Tie
	fmt.Println(stoneGameIII([]int{-1, -2, -3}))  // Tie
}
