package main

import "fmt"

func pathsWithMaxScore(board []string) []int {
	const MOD = 1000000007

	n := len(board)

	// dpScore[i][j] = maximum score from (i,j) to S
	// dpWays[i][j] = number of ways achieving that score
	dpScore := make([][]int, n)
	dpWays := make([][]int, n)

	for i := 0; i < n; i++ {
		dpScore[i] = make([]int, n)
		dpWays[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dpScore[i][j] = -1
		}
	}

	// Starting point: S
	dpScore[n-1][n-1] = 0
	dpWays[n-1][n-1] = 1

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {

			if board[i][j] == 'X' {
				continue
			}

			if i == n-1 && j == n-1 {
				continue
			}

			bestScore := -1
			ways := 0

			// From below
			if i+1 < n && dpWays[i+1][j] > 0 {
				if dpScore[i+1][j] > bestScore {
					bestScore = dpScore[i+1][j]
					ways = dpWays[i+1][j]
				} else if dpScore[i+1][j] == bestScore {
					ways = (ways + dpWays[i+1][j]) % MOD
				}
			}

			// From right
			if j+1 < n && dpWays[i][j+1] > 0 {
				if dpScore[i][j+1] > bestScore {
					bestScore = dpScore[i][j+1]
					ways = dpWays[i][j+1]
				} else if dpScore[i][j+1] == bestScore {
					ways = (ways + dpWays[i][j+1]) % MOD
				}
			}

			// From diagonal (bottom-right)
			if i+1 < n && j+1 < n && dpWays[i+1][j+1] > 0 {
				if dpScore[i+1][j+1] > bestScore {
					bestScore = dpScore[i+1][j+1]
					ways = dpWays[i+1][j+1]
				} else if dpScore[i+1][j+1] == bestScore {
					ways = (ways + dpWays[i+1][j+1]) % MOD
				}
			}

			if ways == 0 {
				continue
			}

			dpScore[i][j] = bestScore
			dpWays[i][j] = ways

			ch := board[i][j]
			if ch >= '1' && ch <= '9' {
				dpScore[i][j] += int(ch - '0')
			}
		}
	}

	if dpWays[0][0] == 0 {
		return []int{0, 0}
	}

	return []int{dpScore[0][0], dpWays[0][0]}
}

func main() {
	tests := [][]string{
		{"E23", "2X2", "12S"},
		{"E12", "1X1", "21S"},
		{"E11", "XXX", "11S"},
	}

	for _, board := range tests {
		fmt.Println(pathsWithMaxScore(board))
	}
}
