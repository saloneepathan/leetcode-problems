package main

import "fmt"

type Pair struct {
	x, y int
}

func maximumSafenessFactor(grid [][]int) int {
	n := len(grid)

	// Step 1: Multi-source BFS to compute distance from nearest thief.
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = -1
		}
	}

	queue := make([]Pair, 0)

	// Push all thief cells into queue.
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				dist[i][j] = 0
				queue = append(queue, Pair{i, j})
			}
		}
	}

	dirs := []int{-1, 0, 1, 0, -1}
	head := 0

	// Multi-source BFS
	for head < len(queue) {
		cur := queue[head]
		head++

		for k := 0; k < 4; k++ {
			nx := cur.x + dirs[k]
			ny := cur.y + dirs[k+1]

			if nx >= 0 && nx < n &&
				ny >= 0 && ny < n &&
				dist[nx][ny] == -1 {

				dist[nx][ny] = dist[cur.x][cur.y] + 1
				queue = append(queue, Pair{nx, ny})
			}
		}
	}

	// BFS to check if a path exists with safeness >= limit.
	canReach := func(limit int) bool {
		if dist[0][0] < limit {
			return false
		}

		visited := make([][]bool, n)
		for i := range visited {
			visited[i] = make([]bool, n)
		}

		q := []Pair{{0, 0}}
		visited[0][0] = true
		head := 0

		for head < len(q) {
			cur := q[head]
			head++

			if cur.x == n-1 && cur.y == n-1 {
				return true
			}

			for k := 0; k < 4; k++ {
				nx := cur.x + dirs[k]
				ny := cur.y + dirs[k+1]

				if nx >= 0 && nx < n &&
					ny >= 0 && ny < n &&
					!visited[nx][ny] &&
					dist[nx][ny] >= limit {

					visited[nx][ny] = true
					q = append(q, Pair{nx, ny})
				}
			}
		}

		return false
	}

	// Binary Search on answer.
	low, high := 0, 2*n
	ans := 0

	for low <= high {
		mid := (low + high) / 2

		if canReach(mid) {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return ans
}

func main() {
	grid1 := [][]int{
		{1, 0, 0},
		{0, 0, 0},
		{0, 0, 1},
	}
	fmt.Println(maximumSafenessFactor(grid1)) // 0

	grid2 := [][]int{
		{0, 0, 1},
		{0, 0, 0},
		{0, 0, 0},
	}
	fmt.Println(maximumSafenessFactor(grid2)) // 2

	grid3 := [][]int{
		{0, 0, 0, 1},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{1, 0, 0, 0},
	}
	fmt.Println(maximumSafenessFactor(grid3)) // 2
}
