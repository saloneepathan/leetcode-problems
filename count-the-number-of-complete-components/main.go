package main

import "fmt"

func countCompleteComponents(n int, edges [][]int) int {
	// Build adjacency list
	graph := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	visited := make([]bool, n)
	completeComponents := 0

	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}

		// DFS to find all vertices in the current component
		stack := []int{i}
		visited[i] = true
		component := []int{}

		for len(stack) > 0 {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			component = append(component, node)

			for _, nei := range graph[node] {
				if !visited[nei] {
					visited[nei] = true
					stack = append(stack, nei)
				}
			}
		}

		// Check if the component is complete
		size := len(component)
		isComplete := true
		for _, node := range component {
			if len(graph[node]) != size-1 {
				isComplete = false
				break
			}
		}

		if isComplete {
			completeComponents++
		}
	}

	return completeComponents
}

func main() {
	// Example 1
	n1 := 6
	edges1 := [][]int{
		{0, 1},
		{0, 2},
		{1, 2},
		{3, 4},
	}
	fmt.Println(countCompleteComponents(n1, edges1)) // Output: 3

	// Example 2
	n2 := 6
	edges2 := [][]int{
		{0, 1},
		{0, 2},
		{1, 2},
		{3, 4},
		{3, 5},
	}
	fmt.Println(countCompleteComponents(n2, edges2)) // Output: 1

	// Additional Test 1
	n3 := 1
	edges3 := [][]int{}
	fmt.Println(countCompleteComponents(n3, edges3)) // Output: 1

	// Additional Test 2
	n4 := 4
	edges4 := [][]int{
		{0, 1},
		{2, 3},
	}
	fmt.Println(countCompleteComponents(n4, edges4)) // Output: 2

	// Additional Test 3
	n5 := 5
	edges5 := [][]int{
		{0, 1},
		{1, 2},
		{2, 0},
		{3, 4},
	}
	fmt.Println(countCompleteComponents(n5, edges5)) // Output: 2
}
