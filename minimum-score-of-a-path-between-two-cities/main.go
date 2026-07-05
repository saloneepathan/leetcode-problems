package main

import "fmt"

func minScore(n int, roads [][]int) int {
	// Build adjacency list: graph[node] = {neighbor, distance}
	graph := make([][][]int, n+1)
	for _, road := range roads {
		u, v, d := road[0], road[1], road[2]
		graph[u] = append(graph[u], []int{v, d})
		graph[v] = append(graph[v], []int{u, d})
	}

	visited := make([]bool, n+1)
	ans := int(^uint(0) >> 1) // Max int

	var dfs func(int)
	dfs = func(node int) {
		visited[node] = true

		for _, edge := range graph[node] {
			next := edge[0]
			dist := edge[1]

			if dist < ans {
				ans = dist
			}

			if !visited[next] {
				dfs(next)
			}
		}
	}

	dfs(1)
	return ans
}

func main() {
	// Example 1
	n1 := 4
	roads1 := [][]int{
		{1, 2, 9},
		{2, 3, 6},
		{2, 4, 5},
		{1, 4, 7},
	}
	fmt.Println(minScore(n1, roads1)) // Output: 5

	// Example 2
	n2 := 4
	roads2 := [][]int{
		{1, 2, 2},
		{1, 3, 4},
		{3, 4, 7},
	}
	fmt.Println(minScore(n2, roads2)) // Output: 2
}
