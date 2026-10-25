
package main

import "fmt"

func closedIsland(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])

	directions := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	var dfs func(int, int) bool
	dfs = func(r, c int) bool {
		// Out of bounds means the island touches the boundary.
		if r < 0 || r >= rows || c < 0 || c >= cols {
			return false
		}

		// Water or already visited cell.
		if grid[r][c] == 1 {
			return true
		}

		// Mark the land cell as visited.
		grid[r][c] = 1

		closed := true

		for _, d := range directions {
			if !dfs(r+d[0], c+d[1]) {
				closed = false
			}
		}

		return closed
	}

	count := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 0 {
				if dfs(r, c) {
					count++
				}
			}
		}
	}

	return count
}

func main() {
	grid1 := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 1, 0},
		{1, 0, 1, 0, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 0},
	}

	grid2 := [][]int{
		{0, 0, 1, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 1, 1, 1, 0},
	}

	grid3 := [][]int{
		{1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 1, 0, 1, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1},
	}

	fmt.Println("Example 1:", closedIsland(grid1))
	fmt.Println("Example 2:", closedIsland(grid2))
	fmt.Println("Example 3:", closedIsland(grid3))
}
