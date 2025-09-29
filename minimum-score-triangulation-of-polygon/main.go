package main

import (
	"fmt"
	"math"
)

func minScoreTriangulation(values []int) int {
	n := len(values)
	// dp[i][j] = minimum score for polygon from i to j
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// length of interval
	for length := 2; length < n; length++ {
		for i := 0; i+length < n; i++ {
			j := i + length
			dp[i][j] = math.MaxInt32
			for k := i + 1; k < j; k++ {
				score := dp[i][k] + dp[k][j] + values[i]*values[j]*values[k]
				if score < dp[i][j] {
					dp[i][j] = score
				}
			}
		}
	}
	return dp[0][n-1]
}

func main() {
	fmt.Println(minScoreTriangulation([]int{1, 2, 3}))          // 6
	fmt.Println(minScoreTriangulation([]int{3, 7, 4, 5}))       // 144
	fmt.Println(minScoreTriangulation([]int{1, 3, 1, 4, 1, 5})) // 13
}
