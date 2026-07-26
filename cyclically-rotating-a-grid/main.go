package main

import "fmt"

func rotateGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])

	layers := min(m, n) / 2

	for layer := 0; layer < layers; layer++ {

		top, left := layer, layer
		bottom, right := m-layer-1, n-layer-1

		// Extract current layer
		var arr []int

		// top row
		for j := left; j <= right; j++ {
			arr = append(arr, grid[top][j])
		}

		// right column
		for i := top + 1; i <= bottom-1; i++ {
			arr = append(arr, grid[i][right])
		}

		// bottom row
		for j := right; j >= left; j-- {
			arr = append(arr, grid[bottom][j])
		}

		// left column
		for i := bottom - 1; i >= top+1; i-- {
			arr = append(arr, grid[i][left])
		}

		size := len(arr)
		rot := k % size

		// Counter-clockwise rotation
		rotated := append(arr[rot:], arr[:rot]...)

		idx := 0

		// Fill top row
		for j := left; j <= right; j++ {
			grid[top][j] = rotated[idx]
			idx++
		}

		// Fill right column
		for i := top + 1; i <= bottom-1; i++ {
			grid[i][right] = rotated[idx]
			idx++
		}

		// Fill bottom row
		for j := right; j >= left; j-- {
			grid[bottom][j] = rotated[idx]
			idx++
		}

		// Fill left column
		for i := bottom - 1; i >= top+1; i-- {
			grid[i][left] = rotated[idx]
			idx++
		}
	}

	return grid
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func printGrid(grid [][]int) {
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println()
}

func main() {

	grid1 := [][]int{
		{40, 10},
		{30, 20},
	}

	k1 := 1

	fmt.Println("Example 1 Output:")
	ans1 := rotateGrid(grid1, k1)
	printGrid(ans1)

	grid2 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	k2 := 2

	fmt.Println("Example 2 Output:")
	ans2 := rotateGrid(grid2, k2)
	printGrid(ans2)
}
