package main

import (
	"fmt"
)

func numMagicSquaresInside(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	// Helper function to check if a 3x3 grid is magic
	isMagic := func(r, c int) bool {
		// Center must be 5
		if grid[r+1][c+1] != 5 {
			return false
		}

		seen := make(map[int]bool)

		// Check numbers are 1-9 and distinct
		for i := r; i < r+3; i++ {
			for j := c; j < c+3; j++ {
				val := grid[i][j]
				if val < 1 || val > 9 || seen[val] {
					return false
				}
				seen[val] = true
			}
		}

		// Sum of first row
		sum := grid[r][c] + grid[r][c+1] + grid[r][c+2]

		// Check rows
		for i := r; i < r+3; i++ {
			if grid[i][c]+grid[i][c+1]+grid[i][c+2] != sum {
				return false
			}
		}

		// Check columns
		for j := c; j < c+3; j++ {
			if grid[r][j]+grid[r+1][j]+grid[r+2][j] != sum {
				return false
			}
		}

		// Check diagonals
		if grid[r][c]+grid[r+1][c+1]+grid[r+2][c+2] != sum {
			return false
		}
		if grid[r][c+2]+grid[r+1][c+1]+grid[r+2][c] != sum {
			return false
		}

		return true
	}

	// Iterate over all possible 3x3 subgrids
	for i := 0; i <= rows-3; i++ {
		for j := 0; j <= cols-3; j++ {
			if isMagic(i, j) {
				count++
			}
		}
	}

	return count
}

func main() {
	// Example 1
	grid1 := [][]int{
		{4, 3, 8, 4},
		{9, 5, 1, 9},
		{2, 7, 6, 2},
	}
	fmt.Println(numMagicSquaresInside(grid1)) // Output: 1

	// Example 2
	grid2 := [][]int{
		{8},
	}
	fmt.Println(numMagicSquaresInside(grid2)) // Output: 0
}
