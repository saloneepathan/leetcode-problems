package main

import (
	"container/heap"
	"fmt"
)

type Cell struct {
	elev int
	x, y int
}

type MinHeap []Cell

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].elev < h[j].elev }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Cell))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func swimInWater(grid [][]int) int {
	n := len(grid)
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	h := &MinHeap{{grid[0][0], 0, 0}}
	heap.Init(h)
	visited[0][0] = true
	res := 0

	for h.Len() > 0 {
		curr := heap.Pop(h).(Cell)
		res = max(res, curr.elev)
		if curr.x == n-1 && curr.y == n-1 {
			return res
		}

		for _, d := range dirs {
			nx, ny := curr.x+d[0], curr.y+d[1]
			if nx >= 0 && ny >= 0 && nx < n && ny < n && !visited[nx][ny] {
				visited[nx][ny] = true
				heap.Push(h, Cell{grid[nx][ny], nx, ny})
			}
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	grid1 := [][]int{{0, 2}, {1, 3}}
	fmt.Println(swimInWater(grid1)) // Output: 3

	grid2 := [][]int{
		{0, 1, 2, 3, 4},
		{24, 23, 22, 21, 5},
		{12, 13, 14, 15, 16},
		{11, 17, 18, 19, 20},
		{10, 9, 8, 7, 6},
	}
	fmt.Println(swimInWater(grid2)) // Output: 16
}
