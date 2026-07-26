package main

import "fmt"

// countUnguarded returns the number of unguarded and unoccupied cells in the grid.
func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	// 0 = empty, 1 = guard, 2 = wall, 3 = guarded
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}

	// Mark guards and walls
	for _, g := range guards {
		grid[g[0]][g[1]] = 1
	}
	for _, w := range walls {
		grid[w[0]][w[1]] = 2
	}

	// Sweep each row (left-right and right-left)
	for i := 0; i < m; i++ {
		seen := false
		// Left to right
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				seen = true
			} else if grid[i][j] == 2 {
				seen = false
			} else if seen {
				grid[i][j] = 3
			}
		}
		// Right to left
		seen = false
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] == 1 {
				seen = true
			} else if grid[i][j] == 2 {
				seen = false
			} else if seen {
				grid[i][j] = 3
			}
		}
	}

	// Sweep each column (top-bottom and bottom-top)
	for j := 0; j < n; j++ {
		seen := false
		// Top to bottom
		for i := 0; i < m; i++ {
			if grid[i][j] == 1 {
				seen = true
			} else if grid[i][j] == 2 {
				seen = false
			} else if seen {
				grid[i][j] = 3
			}
		}
		// Bottom to top
		seen = false
		for i := m - 1; i >= 0; i-- {
			if grid[i][j] == 1 {
				seen = true
			} else if grid[i][j] == 2 {
				seen = false
			} else if seen {
				grid[i][j] = 3
			}
		}
	}

	// Count unguarded (and unoccupied) cells
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				count++
			}
		}
	}

	return count
}

func main() {
	// Example 1
	m1, n1 := 4, 6
	guards1 := [][]int{{0, 0}, {1, 1}, {2, 3}}
	walls1 := [][]int{{0, 1}, {2, 2}, {1, 4}}
	fmt.Println("Example 1 Output:", countUnguarded(m1, n1, guards1, walls1)) // Expected: 7

	// Example 2
	m2, n2 := 3, 3
	guards2 := [][]int{{1, 1}}
	walls2 := [][]int{{0, 1}, {1, 0}, {2, 1}, {1, 2}}
	fmt.Println("Example 2 Output:", countUnguarded(m2, n2, guards2, walls2)) // Expected: 4
}
