package main

import (
	"fmt"
	"math"
	"sort"
)

func minCost(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])

	type point struct{ x, y int }

	// Collect all points
	points := make([]point, 0, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			points = append(points, point{i, j})
		}
	}

	// Sort points by cell value
	sort.Slice(points, func(i, j int) bool {
		return grid[points[i].x][points[i].y] <
			grid[points[j].x][points[j].y]
	})

	// costs[i][j] = min cost from (i,j) to bottom-right
	costs := make([][]int, m)
	for i := range costs {
		costs[i] = make([]int, n)
		for j := range costs[i] {
			costs[i][j] = math.MaxInt
		}
	}

	// Each iteration allows one more teleport
	for t := 0; t <= k; t++ {
		// Teleport relaxation (value sweep)
		minCostSoFar := math.MaxInt
		start := 0

		for i := 0; i < len(points); i++ {
			x, y := points[i].x, points[i].y
			minCostSoFar = min(minCostSoFar, costs[x][y])

			// Process equal-value group together
			if i+1 < len(points) &&
				grid[points[i].x][points[i].y] ==
					grid[points[i+1].x][points[i+1].y] {
				continue
			}

			for r := start; r <= i; r++ {
				px, py := points[r].x, points[r].y
				costs[px][py] = min(costs[px][py], minCostSoFar)
			}
			start = i + 1
		}

		// Normal DP moves (right & down)
		for i := m - 1; i >= 0; i-- {
			for j := n - 1; j >= 0; j-- {
				if i == m-1 && j == n-1 {
					costs[i][j] = 0
					continue
				}
				if i+1 < m {
					costs[i][j] = min(costs[i][j], costs[i+1][j]+grid[i+1][j])
				}
				if j+1 < n {
					costs[i][j] = min(costs[i][j], costs[i][j+1]+grid[i][j+1])
				}
			}
		}
	}

	return costs[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	grid1 := [][]int{
		{1, 3, 3},
		{2, 5, 4},
		{4, 3, 5},
	}
	k1 := 2
	fmt.Println("Output:", minCost(grid1, k1)) // Expected: 7

	grid2 := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
	}
	k2 := 1
	fmt.Println("Output:", minCost(grid2, k2)) // Expected: 9
}
