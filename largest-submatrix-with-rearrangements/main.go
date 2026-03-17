package main

import (
	"fmt"
	"sort"
)

func largestSubmatrix(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])

	// Build heights (histogram)
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				matrix[i][j] += matrix[i-1][j]
			}
		}
	}

	maxArea := 0

	// Process each row
	for i := 0; i < m; i++ {
		row := make([]int, n)
		copy(row, matrix[i])

		// Sort descending
		sort.Slice(row, func(a, b int) bool {
			return row[a] > row[b]
		})

		// Calculate max area
		for j := 0; j < n; j++ {
			area := row[j] * (j + 1)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func main() {
	matrix1 := [][]int{
		{0, 0, 1},
		{1, 1, 1},
		{1, 0, 1},
	}

	matrix2 := [][]int{
		{1, 0, 1, 0, 1},
	}

	matrix3 := [][]int{
		{1, 1, 0},
		{1, 0, 1},
	}

	fmt.Println("Output 1:", largestSubmatrix(matrix1)) // 4
	fmt.Println("Output 2:", largestSubmatrix(matrix2)) // 3
	fmt.Println("Output 3:", largestSubmatrix(matrix3)) // 2
}
