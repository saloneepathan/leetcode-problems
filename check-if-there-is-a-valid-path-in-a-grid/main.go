package main

import "fmt"

func hasValidPath(grid [][]int) bool {
	m, n := len(grid), len(grid[0])

	// Directions: up, down, left, right
	dirs := [][]int{
		{-1, 0}, // up
		{1, 0},  // down
		{0, -1}, // left
		{0, 1},  // right
	}

	// Street type to allowed directions
	streetDirs := map[int][]int{
		1: {2, 3}, // left, right
		2: {0, 1}, // up, down
		3: {2, 1}, // left, down
		4: {3, 1}, // right, down
		5: {2, 0}, // left, up
		6: {3, 0}, // right, up
	}

	// Opposite directions
	opposite := map[int]int{
		0: 1, // up -> down
		1: 0, // down -> up
		2: 3, // left -> right
		3: 2, // right -> left
	}

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	queue := [][]int{{0, 0}}
	visited[0][0] = true

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]

		x, y := cell[0], cell[1]

		if x == m-1 && y == n-1 {
			return true
		}

		for _, dir := range streetDirs[grid[x][y]] {
			nx := x + dirs[dir][0]
			ny := y + dirs[dir][1]

			if nx < 0 || ny < 0 || nx >= m || ny >= n || visited[nx][ny] {
				continue
			}

			// Check if neighbor connects back
			for _, backDir := range streetDirs[grid[nx][ny]] {
				if backDir == opposite[dir] {
					visited[nx][ny] = true
					queue = append(queue, []int{nx, ny})
					break
				}
			}
		}
	}

	return false
}

func main() {
	grid1 := [][]int{
		{2, 4, 3},
		{6, 5, 2},
	}
	fmt.Println(hasValidPath(grid1)) // true

	grid2 := [][]int{
		{1, 2, 1},
		{1, 2, 1},
	}
	fmt.Println(hasValidPath(grid2)) // false

	grid3 := [][]int{
		{1, 1, 2},
	}
	fmt.Println(hasValidPath(grid3)) // false
}
