package main

import "fmt"

func canPartitionGrid(grid [][]int) bool {
	m := len(grid)
	n := len(grid[0])

	total := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			total += grid[i][j]
		}
	}

	// if total is odd, cannot split
	if total%2 != 0 {
		return false
	}

	target := total / 2

	// Try horizontal cut
	rowSum := 0
	for i := 0; i < m-1; i++ { // ensure non-empty bottom part
		for j := 0; j < n; j++ {
			rowSum += grid[i][j]
		}
		if rowSum == target {
			return true
		}
	}

	// Try vertical cut
	colSum := make([]int, n)
	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			colSum[j] += grid[i][j]
		}
	}

	sum := 0
	for j := 0; j < n-1; j++ { // ensure non-empty right part
		sum += colSum[j]
		if sum == target {
			return true
		}
	}

	return false
}

func main() {
	grid1 := [][]int{{1, 4}, {2, 3}}
	fmt.Println(canPartitionGrid(grid1)) // true

	grid2 := [][]int{{1, 3}, {2, 4}}
	fmt.Println(canPartitionGrid(grid2)) // false
}
