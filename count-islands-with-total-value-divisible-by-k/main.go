package main

import "fmt"

func countIslands(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])

	directions := [][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	var dfs func(int, int) int
	dfs = func(r, c int) int {
		if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] == 0 {
			return 0
		}

		sum := grid[r][c]
		grid[r][c] = 0 // Mark as visited

		for _, d := range directions {
			sum += dfs(r+d[0], c+d[1])
		}

		return sum
	}

	count := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				islandSum := dfs(i, j)

				if islandSum%k == 0 {
					count++
				}
			}
		}
	}

	return count
}

func main() {
	grid1 := [][]int{
		{0, 2, 1, 0, 0},
		{0, 5, 0, 0, 5},
		{0, 0, 1, 0, 0},
		{0, 1, 4, 7, 0},
		{0, 2, 0, 0, 8},
	}
	k1 := 5
	fmt.Println(countIslands(grid1, k1)) // Output: 2

	grid2 := [][]int{
		{3, 0, 3, 0},
		{0, 3, 0, 3},
		{3, 0, 3, 0},
	}
	k2 := 3
	fmt.Println(countIslands(grid2, k2)) // Output: 6
}
