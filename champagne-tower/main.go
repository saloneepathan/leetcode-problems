package main

import (
	"fmt"
	"math"
)

func champagneTower(poured int, query_row int, query_glass int) float64 {

	// create DP array
	dp := make([][]float64, query_row+2)

	for i := range dp {
		dp[i] = make([]float64, query_row+2)
	}

	dp[0][0] = float64(poured)

	for i := 0; i <= query_row; i++ {
		for j := 0; j <= i; j++ {

			if dp[i][j] > 1 {

				overflow := (dp[i][j] - 1) / 2

				dp[i+1][j] += overflow
				dp[i+1][j+1] += overflow
			}
		}
	}

	return math.Min(1, dp[query_row][query_glass])
}

func main() {

	fmt.Println(champagneTower(1, 1, 1)) // 0.0
	fmt.Println(champagneTower(2, 1, 1)) // 0.5
	fmt.Println(champagneTower(4, 2, 1)) // 0.5
}
