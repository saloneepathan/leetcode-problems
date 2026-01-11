package main

import "fmt"

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	rows, cols := len(matrix), len(matrix[0])
	heights := make([]int, cols)
	maxArea := 0

	for i := 0; i < rows; i++ {
		// Build histogram heights for current row
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}
		maxArea = max(maxArea, largestRectangleArea(heights))
	}

	return maxArea
}

func largestRectangleArea(heights []int) int {
	stack := []int{}
	maxArea := 0

	// Add sentinel to flush stack at end
	heights = append(heights, 0)

	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}

			maxArea = max(maxArea, h*width)
		}
		stack = append(stack, i)
	}

	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test Case 1
	matrix1 := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println("Output:", maximalRectangle(matrix1)) // Expected: 6

	// Test Case 2
	matrix2 := [][]byte{
		{'0'},
	}
	fmt.Println("Output:", maximalRectangle(matrix2)) // Expected: 0

	// Test Case 3
	matrix3 := [][]byte{
		{'1'},
	}
	fmt.Println("Output:", maximalRectangle(matrix3)) // Expected: 1
}
