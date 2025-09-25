package main

import (
	"fmt"
	"math"
)

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([]int, n)

	// Initialize dp with the last row
	for i := 0; i < n; i++ {
		dp[i] = triangle[n-1][i]
	}

	// Bottom-up DP
	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = triangle[i][j] + int(math.Min(float64(dp[j]), float64(dp[j+1])))
		}
	}

	return dp[0]
}

func main() {
	triangle1 := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	triangle2 := [][]int{{-10}}

	fmt.Println(minimumTotal(triangle1)) // Output: 11
	fmt.Println(minimumTotal(triangle2)) // Output: -10
}
