package main

import "fmt"

func findSafeWalk(grid [][]int, health int) bool {
	m, n := len(grid), len(grid[0])

	const INF = int(1e9)

	// dist[i][j] = minimum health lost to reach (i,j)
	dist := make([][]int, m)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = INF
		}
	}

	type Node struct {
		x, y int
	}

	// Deque implementation
	deque := make([]Node, 0)

	pushFront := func(v Node) {
		deque = append([]Node{v}, deque...)
	}

	pushBack := func(v Node) {
		deque = append(deque, v)
	}

	popFront := func() Node {
		v := deque[0]
		deque = deque[1:]
		return v
	}

	dist[0][0] = grid[0][0]
	pushFront(Node{0, 0})

	dirs := []int{-1, 0, 1, 0, -1}

	for len(deque) > 0 {
		cur := popFront()

		for k := 0; k < 4; k++ {
			nx := cur.x + dirs[k]
			ny := cur.y + dirs[k+1]

			if nx < 0 || nx >= m || ny < 0 || ny >= n {
				continue
			}

			newCost := dist[cur.x][cur.y] + grid[nx][ny]

			if newCost < dist[nx][ny] {
				dist[nx][ny] = newCost

				if grid[nx][ny] == 0 {
					pushFront(Node{nx, ny})
				} else {
					pushBack(Node{nx, ny})
				}
			}
		}
	}

	return dist[m-1][n-1] < health
}

func main() {
	grid1 := [][]int{
		{0, 1, 0, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 1, 0},
	}
	fmt.Println(findSafeWalk(grid1, 1)) // true

	grid2 := [][]int{
		{0, 1, 1, 0, 0, 0},
		{1, 0, 1, 0, 0, 0},
		{0, 1, 1, 1, 0, 1},
		{0, 0, 1, 0, 1, 0},
	}
	fmt.Println(findSafeWalk(grid2, 3)) // false

	grid3 := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	fmt.Println(findSafeWalk(grid3, 5)) // true
}
