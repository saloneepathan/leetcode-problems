package main

import "fmt"

func stoneGame(piles []int) bool {
	n := len(piles)

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = piles[i]
	}

	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1

			left := piles[i] - dp[i+1][j]
			right := piles[j] - dp[i][j-1]

			if left > right {
				dp[i][j] = left
			} else {
				dp[i][j] = right
			}
		}
	}

	return dp[0][n-1] > 0
}

func main() {
	fmt.Println(stoneGame([]int{5, 3, 4, 5})) // true
	fmt.Println(stoneGame([]int{3, 7, 2, 3})) // true
}