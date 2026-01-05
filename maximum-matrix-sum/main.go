package main

import (
	"fmt"
)

func maxMatrixSum(matrix [][]int) int64 {
	var sum int64 = 0
	negCount := 0
	minAbs := int64(1e18)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			val := int64(matrix[i][j])
			if val < 0 {
				negCount++
				val = -val
			}
			sum += val
			if val < minAbs {
				minAbs = val
			}
		}
	}

	if negCount%2 == 1 {
		sum -= 2 * minAbs
	}

	return sum
}

func main() {
	// Example 1
	matrix1 := [][]int{
		{1, -1},
		{-1, 1},
	}
	fmt.Println(maxMatrixSum(matrix1)) // Output: 4

	// Example 2
	matrix2 := [][]int{
		{1, 2, 3},
		{-1, -2, -3},
		{1, 2, 3},
	}
	fmt.Println(maxMatrixSum(matrix2)) // Output: 16
}
