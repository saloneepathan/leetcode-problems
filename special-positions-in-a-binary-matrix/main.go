package main

import (
	"fmt"
)

func numSpecial(mat [][]int) int {
	m := len(mat)
	n := len(mat[0])

	rowCount := make([]int, m)
	colCount := make([]int, n)

	// First pass: count 1s in each row and column
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				rowCount[i]++
				colCount[j]++
			}
		}
	}

	special := 0

	// Second pass: check special positions
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 && rowCount[i] == 1 && colCount[j] == 1 {
				special++
			}
		}
	}

	return special
}

func main() {
	mat1 := [][]int{
		{1, 0, 0},
		{0, 0, 1},
		{1, 0, 0},
	}

	mat2 := [][]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}

	fmt.Println("Output 1:", numSpecial(mat1)) // Expected: 1
	fmt.Println("Output 2:", numSpecial(mat2)) // Expected: 3
}
