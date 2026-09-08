package main

import "fmt"

func winnerSquareGame(n int) bool {
	memo := make([]int, n+1)

	// 0 = not calculated
	// 1 = false
	// 2 = true
	var dfs func(int) bool

	dfs = func(i int) bool {
		if i == 0 {
			return false
		}

		if memo[i] != 0 {
			return memo[i] == 2
		}

		for j := 1; j*j <= i; j++ {
			// If there is a move that leaves the opponent
			// in a losing state, current player wins.
			if !dfs(i - j*j) {
				memo[i] = 2
				return true
			}
		}

		memo[i] = 1
		return false
	}

	return dfs(n)
}

func main() {
	n := 7

	result := winnerSquareGame(n)

	fmt.Println(result)
}
