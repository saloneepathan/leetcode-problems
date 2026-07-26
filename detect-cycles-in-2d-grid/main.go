package main

import "fmt"

func containsCycle(grid [][]byte) bool {
	m, n := len(grid), len(grid[0])

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	dirs := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	var dfs func(x, y, px, py int, ch byte) bool
	dfs = func(x, y, px, py int, ch byte) bool {
		visited[x][y] = true

		for _, d := range dirs {
			nx, ny := x+d[0], y+d[1]

			// Boundary check
			if nx < 0 || ny < 0 || nx >= m || ny >= n {
				continue
			}

			// Only move to same character cells
			if grid[nx][ny] != ch {
				continue
			}

			// Skip the previous cell
			if nx == px && ny == py {
				continue
			}

			// Cycle detected
			if visited[nx][ny] {
				return true
			}

			if dfs(nx, ny, x, y, ch) {
				return true
			}
		}

		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !visited[i][j] {
				if dfs(i, j, -1, -1, grid[i][j]) {
					return true
				}
			}
		}
	}

	return false
}

func main() {
	grid1 := [][]byte{
		{'a', 'a', 'a', 'a'},
		{'a', 'b', 'b', 'a'},
		{'a', 'b', 'b', 'a'},
		{'a', 'a', 'a', 'a'},
	}

	grid2 := [][]byte{
		{'c', 'c', 'c', 'a'},
		{'c', 'd', 'c', 'c'},
		{'c', 'c', 'e', 'c'},
		{'f', 'c', 'c', 'c'},
	}

	grid3 := [][]byte{
		{'a', 'b', 'b'},
		{'b', 'z', 'b'},
		{'b', 'b', 'a'},
	}

	fmt.Println(containsCycle(grid1)) // true
	fmt.Println(containsCycle(grid2)) // true
	fmt.Println(containsCycle(grid3)) // false
}
