package main

import "fmt"

func largestMagicSquare(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// Prefix sums
	row := make([][]int, m)
	col := make([][]int, m+1)
	diag1 := make([][]int, m+1) // main diagonal ↘
	diag2 := make([][]int, m+1) // anti diagonal ↙

	for i := 0; i < m; i++ {
		row[i] = make([]int, n+1)
		col[i] = make([]int, n)
		diag1[i] = make([]int, n+1)
		diag2[i] = make([]int, n+2)
	}
	col[m] = make([]int, n)
	diag1[m] = make([]int, n+1)
	diag2[m] = make([]int, n+2)

	// Build prefix sums
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			row[i][j+1] = row[i][j] + grid[i][j]
			col[i+1][j] = col[i][j] + grid[i][j]
			diag1[i+1][j+1] = diag1[i][j] + grid[i][j]
			diag2[i+1][j] = diag2[i][j+1] + grid[i][j]
		}
	}

	maxSize := min(m, n)

	// Try larger squares first
	for size := maxSize; size >= 2; size-- {
		for r := 0; r+size <= m; r++ {
			for c := 0; c+size <= n; c++ {

				target := row[r][c+size] - row[r][c]
				valid := true

				// Check rows
				for i := 0; i < size && valid; i++ {
					if row[r+i][c+size]-row[r+i][c] != target {
						valid = false
					}
				}

				// Check columns
				for j := 0; j < size && valid; j++ {
					if col[r+size][c+j]-col[r][c+j] != target {
						valid = false
					}
				}

				// Check diagonals
				d1 := diag1[r+size][c+size] - diag1[r][c]
				d2 := diag2[r+size][c] - diag2[r][c+size]
				if d1 != target || d2 != target {
					valid = false
				}

				if valid {
					return size
				}
			}
		}
	}

	return 1 // every 1x1 grid is a magic square
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	grid1 := [][]int{
		{7, 1, 4, 5, 6},
		{2, 5, 1, 6, 4},
		{1, 5, 4, 3, 2},
		{1, 2, 7, 3, 4},
	}
	fmt.Println(largestMagicSquare(grid1)) // Output: 3

	grid2 := [][]int{
		{5, 1, 3, 1},
		{9, 3, 3, 1},
		{1, 3, 3, 8},
	}
	fmt.Println(largestMagicSquare(grid2)) // Output: 2
}
