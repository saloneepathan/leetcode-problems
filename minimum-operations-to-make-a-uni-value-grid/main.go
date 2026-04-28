package main

import (
	"fmt"
	"sort"
)

func minOperations(grid [][]int, x int) int {
	// Flatten the grid into a single array
	nums := []int{}
	remainder := grid[0][0] % x

	for _, row := range grid {
		for _, val := range row {
			// Check feasibility
			if val%x != remainder {
				return -1
			}
			nums = append(nums, val)
		}
	}

	// Sort to find median
	sort.Ints(nums)

	// Median minimizes total absolute difference
	median := nums[len(nums)/2]

	operations := 0
	for _, val := range nums {
		diff := val - median
		if diff < 0 {
			diff = -diff
		}
		operations += diff / x
	}

	return operations
}

func main() {
	grid1 := [][]int{{2, 4}, {6, 8}}
	x1 := 2
	fmt.Println(minOperations(grid1, x1)) // Output: 4

	grid2 := [][]int{{1, 5}, {2, 3}}
	x2 := 1
	fmt.Println(minOperations(grid2, x2)) // Output: 5

	grid3 := [][]int{{1, 2}, {3, 4}}
	x3 := 2
	fmt.Println(minOperations(grid3, x3)) // Output: -1
}
