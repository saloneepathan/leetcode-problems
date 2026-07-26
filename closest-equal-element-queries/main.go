package main

import (
	"fmt"
	"sort"
)

func solveQueries(nums []int, queries []int) []int {
	n := len(nums)

	// Map value -> indices
	pos := make(map[int][]int)
	for i, v := range nums {
		pos[v] = append(pos[v], i)
	}

	res := make([]int, len(queries))

	for qi, idx := range queries {
		val := nums[idx]
		indices := pos[val]

		// If only one occurrence
		if len(indices) == 1 {
			res[qi] = -1
			continue
		}

		// Binary search
		i := sort.SearchInts(indices, idx)

		best := n

		// Left neighbor
		left := (i - 1 + len(indices)) % len(indices)
		dist := abs(indices[left] - idx)
		best = min(best, min(dist, n-dist))

		// Right neighbor
		right := (i + 1) % len(indices)
		dist = abs(indices[right] - idx)
		best = min(best, min(dist, n-dist))

		res[qi] = best
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 2, 1, 3, 1}
	queries := []int{0, 1, 2, 3, 4}

	result := solveQueries(nums, queries)
	fmt.Println(result) // Expected: [2 -1 2 -1 2]
}
