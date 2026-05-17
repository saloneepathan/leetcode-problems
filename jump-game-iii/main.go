package main

import "fmt"

func canReach(arr []int, start int) bool {
	n := len(arr)
	visited := make([]bool, n)

	var dfs func(int) bool
	dfs = func(i int) bool {
		// Out of bounds or already visited
		if i < 0 || i >= n || visited[i] {
			return false
		}

		// Found zero
		if arr[i] == 0 {
			return true
		}

		visited[i] = true

		// Try both directions
		return dfs(i+arr[i]) || dfs(i-arr[i])
	}

	return dfs(start)
}

func main() {
	arr1 := []int{4, 2, 3, 0, 3, 1, 2}
	start1 := 5
	fmt.Println(canReach(arr1, start1)) // true

	arr2 := []int{4, 2, 3, 0, 3, 1, 2}
	start2 := 0
	fmt.Println(canReach(arr2, start2)) // true

	arr3 := []int{3, 0, 2, 1, 2}
	start3 := 2
	fmt.Println(canReach(arr3, start3)) // false
}
