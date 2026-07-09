package main

import "fmt"

func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {
	parent := make([]int, n)
	rank := make([]int, n)

	// Initialize DSU
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	// Find with path compression
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	// Union by rank
	union := func(x, y int) {
		px, py := find(x), find(y)
		if px == py {
			return
		}

		if rank[px] < rank[py] {
			parent[px] = py
		} else if rank[px] > rank[py] {
			parent[py] = px
		} else {
			parent[py] = px
			rank[px]++
		}
	}

	// Connect adjacent nodes if possible
	for i := 0; i < n-1; i++ {
		if nums[i+1]-nums[i] <= maxDiff {
			union(i, i+1)
		}
	}

	// Answer queries
	ans := make([]bool, len(queries))
	for i, q := range queries {
		ans[i] = find(q[0]) == find(q[1])
	}

	return ans
}

func main() {
	// Example 1
	n1 := 2
	nums1 := []int{1, 3}
	maxDiff1 := 1
	queries1 := [][]int{{0, 0}, {0, 1}}

	fmt.Println(pathExistenceQueries(n1, nums1, maxDiff1, queries1))
	// Output: [true false]

	// Example 2
	n2 := 4
	nums2 := []int{2, 5, 6, 8}
	maxDiff2 := 2
	queries2 := [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}

	fmt.Println(pathExistenceQueries(n2, nums2, maxDiff2, queries2))
	// Output: [false false true true]
}
