package main

import (
	"fmt"
)

type Solution struct{}

func (Solution) maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	// Build adjacency list
	g := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	components := 0

	var dfs func(int, int) int64
	dfs = func(node, parent int) int64 {
		var subtotal int64 = int64(values[node])

		for _, nei := range g[node] {
			if nei == parent {
				continue
			}
			subtotal += dfs(nei, node)
		}

		if subtotal%int64(k) == 0 {
			components++
			return 0
		}
		return subtotal
	}

	dfs(0, -1)
	return components
}

func main() {
	// Test example
	n := 5
	edges := [][]int{{0, 2}, {1, 2}, {1, 3}, {2, 4}}
	values := []int{1, 8, 1, 4, 4}
	k := 6

	sol := Solution{}
	fmt.Println(sol.maxKDivisibleComponents(n, edges, values, k)) // Output: 2
}
