package main

import (
	"fmt"
	"sort"
)

func sortMatrix(grid [][]int) [][]int {
	diagonals := make(map[int][]int)

	// Group elements by diagonal key (i - j)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			key := i - j
			diagonals[key] = append(diagonals[key], grid[i][j])
		}
	}

	// Sort each diagonal: ascending for lower, descending for upper
	for key, arr := range diagonals {
		if key >= 0 {
			sort.Sort(sort.Reverse(sort.IntSlice(arr))) // descending
		} else {
			sort.Ints(arr) // ascending
		}
		diagonals[key] = arr
	}

	// Fill grid back with sorted values
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			key := i - j
			grid[i][j] = diagonals[key][0]
			diagonals[key] = diagonals[key][1:]
		}
	}

	return grid
}

func main() {
	// Example matrix
	grid := [][]int{
		{8, 4, 1},
		{4, 4, 1},
		{4, 8, 9},
	}

	fmt.Println("Original matrix:")
	for _, row := range grid {
		fmt.Println(row)
	}

	// Sort the diagonals
	result := sortMatrix(grid)

	fmt.Println("\nMatrix after sorting diagonals:")
	for _, row := range result {
		fmt.Println(row)
	}
}
