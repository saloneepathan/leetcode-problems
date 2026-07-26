package main

import (
	"fmt"
)

func countSubmatrices(grid [][]int, k int) int {
	n := len(grid)
	m := len(grid[0])

	cols := make([]int, m)
	res := 0

	for i := 0; i < n; i++ {
		rowSum := 0
		for j := 0; j < m; j++ {
			// accumulate column sums
			cols[j] += grid[i][j]

			// accumulate row-wise sum of rectangle (0,0) to (i,j)
			rowSum += cols[j]

			if rowSum <= k {
				res++
			}
		}
	}

	return res
}

func main() {
	// Test Case 1
	grid1 := [][]int{
		{7, 6, 3},
		{6, 6, 1},
	}
	k1 := 18
	fmt.Println("Output 1:", countSubmatrices(grid1, k1)) // Expected: 4

	// Test Case 2
	grid2 := [][]int{
		{7, 2, 9},
		{1, 5, 0},
		{2, 6, 6},
	}
	k2 := 20
	fmt.Println("Output 2:", countSubmatrices(grid2, k2)) // Expected: 6
}
