package main

import "fmt"

func latestDayToCross(row int, col int, cells [][]int) int {
	left, right := 1, row*col
	ans := 0

	for left <= right {
		mid := left + (right-left)/2
		if canCross(row, col, cells, mid) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}

func canCross(row, col int, cells [][]int, day int) bool {
	// grid[r][c] = 0 -> land, 1 -> water
	grid := make([][]int, row)
	for i := 0; i < row; i++ {
		grid[i] = make([]int, col)
	}

	// Flood cells for given day
	for i := 0; i < day; i++ {
		r := cells[i][0] - 1
		c := cells[i][1] - 1
		grid[r][c] = 1
	}

	visited := make([][]bool, row)
	for i := 0; i < row; i++ {
		visited[i] = make([]bool, col)
	}

	queue := [][]int{}

	// Start BFS from top row
	for c := 0; c < col; c++ {
		if grid[0][c] == 0 {
			queue = append(queue, []int{0, c})
			visited[0][c] = true
		}
	}

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		r, c := cell[0], cell[1]

		// Reached bottom row
		if r == row-1 {
			return true
		}

		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < row && nc >= 0 && nc < col {
				if !visited[nr][nc] && grid[nr][nc] == 0 {
					visited[nr][nc] = true
					queue = append(queue, []int{nr, nc})
				}
			}
		}
	}
	return false
}

func main() {
	row := 3
	col := 3
	cells := [][]int{
		{1, 2}, {2, 1}, {3, 3},
		{2, 2}, {1, 1}, {1, 3},
		{2, 3}, {3, 2}, {3, 1},
	}

	result := latestDayToCross(row, col, cells)
	fmt.Println("Last day to cross:", result)
}
