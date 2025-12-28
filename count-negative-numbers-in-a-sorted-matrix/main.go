package main

import "fmt"

func countNegatives(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	row := m - 1
	col := 0
	count := 0

	for row >= 0 && col < n {
		if grid[row][col] < 0 {
			count += n - col
			row--
		} else {
			col++
		}
	}

	return count
}

func main() {
	grid1 := [][]int{
		{4, 3, 2, -1},
		{3, 2, 1, -1},
		{1, 1, -1, -2},
		{-1, -1, -2, -3},
	}

	grid2 := [][]int{
		{3, 2},
		{1, 0},
	}

	fmt.Println(countNegatives(grid1)) // Output: 8
	fmt.Println(countNegatives(grid2)) // Output: 0
}
