package main

import "fmt"

func remainingMethods(n int, k int, invocations [][]int) []int {
	// Build graph
	graph := make([][]int, n)
	for _, edge := range invocations {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
	}

	// Mark suspicious methods using DFS
	suspicious := make([]bool, n)

	var dfs func(int)
	dfs = func(node int) {
		if suspicious[node] {
			return
		}
		suspicious[node] = true
		for _, next := range graph[node] {
			dfs(next)
		}
	}

	dfs(k)

	// Check if any non-suspicious method invokes a suspicious method
	for _, edge := range invocations {
		u, v := edge[0], edge[1]
		if !suspicious[u] && suspicious[v] {
			// Cannot remove any methods
			ans := make([]int, n)
			for i := 0; i < n; i++ {
				ans[i] = i
			}
			return ans
		}
	}

	// Return remaining methods
	ans := []int{}
	for i := 0; i < n; i++ {
		if !suspicious[i] {
			ans = append(ans, i)
		}
	}

	return ans
}

func main() {
	// Example 1
	n1 := 4
	k1 := 1
	invocations1 := [][]int{
		{1, 2},
		{0, 1},
		{3, 2},
	}
	fmt.Println("Example 1:", remainingMethods(n1, k1, invocations1))

	// Example 2
	n2 := 5
	k2 := 0
	invocations2 := [][]int{
		{1, 2},
		{0, 2},
		{0, 1},
		{3, 4},
	}
	fmt.Println("Example 2:", remainingMethods(n2, k2, invocations2))

	// Example 3
	n3 := 3
	k3 := 2
	invocations3 := [][]int{
		{1, 2},
		{0, 1},
		{2, 0},
	}
	fmt.Println("Example 3:", remainingMethods(n3, k3, invocations3))
}
