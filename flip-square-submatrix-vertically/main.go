package main

import "fmt"

func reverseSubmatrix(grid [][]int, x int, y int, k int) [][]int {
	top := x
	bottom := x + k - 1

	for top < bottom {
		for col := y; col < y+k; col++ {
			grid[top][col], grid[bottom][col] = grid[bottom][col], grid[top][col]
		}
		top++
		bottom--
	}

	return grid
}

func printGrid(grid [][]int) {
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println()
}

func main() {
	// Example 1
	grid1 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	x1, y1, k1 := 1, 0, 3

	fmt.Println("Example 1 Output:")
	result1 := reverseSubmatrix(grid1, x1, y1, k1)
	printGrid(result1)

	// Example 2
	grid2 := [][]int{
		{3, 4, 2, 3},
		{2, 3, 4, 2},
	}
	x2, y2, k2 := 0, 2, 2

	fmt.Println("Example 2 Output:")
	result2 := reverseSubmatrix(grid2, x2, y2, k2)
	printGrid(result2)
}
