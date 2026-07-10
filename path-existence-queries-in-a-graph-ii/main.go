package main

import (
	"fmt"
	"sort"
)

func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []int {
	type Node struct {
		val int
		idx int
	}

	// Sort nodes by value
	arr := make([]Node, n)
	for i := 0; i < n; i++ {
		arr[i] = Node{nums[i], i}
	}

	sort.Slice(arr, func(i, j int) bool {
		if arr[i].val == arr[j].val {
			return arr[i].idx < arr[j].idx
		}
		return arr[i].val < arr[j].val
	})

	// Position of each original index in sorted order
	pos := make([]int, n)
	for i := 0; i < n; i++ {
		pos[arr[i].idx] = i
	}

	// Connected component IDs
	comp := make([]int, n)
	component := 0
	comp[0] = component
	for i := 1; i < n; i++ {
		if arr[i].val-arr[i-1].val > maxDiff {
			component++
		}
		comp[i] = component
	}

	// reach[i] = furthest sorted index reachable in one edge
	reach := make([]int, n)
	j := 0
	for i := 0; i < n; i++ {
		if j < i {
			j = i
		}
		for j+1 < n && arr[j+1].val-arr[i].val <= maxDiff {
			j++
		}
		reach[i] = j
	}

	// Binary lifting table
	LOG := 1
	for (1 << LOG) <= n {
		LOG++
	}

	up := make([][]int, LOG)
	up[0] = make([]int, n)
	copy(up[0], reach)

	for k := 1; k < LOG; k++ {
		up[k] = make([]int, n)
		for i := 0; i < n; i++ {
			up[k][i] = up[k-1][up[k-1][i]]
		}
	}

	ans := make([]int, len(queries))

	for qi, q := range queries {
		u := pos[q[0]]
		v := pos[q[1]]

		if u > v {
			u, v = v, u
		}

		// Different connected components
		if comp[u] != comp[v] {
			ans[qi] = -1
			continue
		}

		if u == v {
			ans[qi] = 0
			continue
		}

		cur := u
		steps := 0

		// Binary lifting to count minimum jumps
		for k := LOG - 1; k >= 0; k-- {
			if up[k][cur] < v {
				cur = up[k][cur]
				steps += 1 << k
			}
		}

		ans[qi] = steps + 1
	}

	return ans
}

func main() {
	// Example 1
	n1 := 5
	nums1 := []int{1, 8, 3, 4, 2}
	maxDiff1 := 3
	queries1 := [][]int{{0, 3}, {2, 4}}
	fmt.Println(pathExistenceQueries(n1, nums1, maxDiff1, queries1))
	// Output: [1 1]

	// Example 2
	n2 := 5
	nums2 := []int{5, 3, 1, 9, 10}
	maxDiff2 := 2
	queries2 := [][]int{{0, 1}, {0, 2}, {2, 3}, {4, 3}}
	fmt.Println(pathExistenceQueries(n2, nums2, maxDiff2, queries2))
	// Output: [1 2 -1 1]

	// Example 3
	n3 := 3
	nums3 := []int{3, 6, 1}
	maxDiff3 := 1
	queries3 := [][]int{{0, 0}, {0, 1}, {1, 2}}
	fmt.Println(pathExistenceQueries(n3, nums3, maxDiff3, queries3))
	// Output: [0 -1 -1]
}
