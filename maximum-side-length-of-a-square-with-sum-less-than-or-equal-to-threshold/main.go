package main

import "fmt"

func maxSideLength(mat [][]int, threshold int) int {
	m, n := len(mat), len(mat[0])

	// Build prefix sum matrix
	prefix := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		prefix[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			prefix[i+1][j+1] = mat[i][j] +
				prefix[i][j+1] +
				prefix[i+1][j] -
				prefix[i][j]
		}
	}

	// Check if any k x k square has sum <= threshold
	canForm := func(k int) bool {
		for i := k; i <= m; i++ {
			for j := k; j <= n; j++ {
				sum := prefix[i][j] -
					prefix[i-k][j] -
					prefix[i][j-k] +
					prefix[i-k][j-k]
				if sum <= threshold {
					return true
				}
			}
		}
		return false
	}

	left, right := 0, min(m, n)
	ans := 0

	// Binary search on side length
	for left <= right {
		mid := (left + right) / 2
		if canForm(mid) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	mat1 := [][]int{
		{1, 1, 3, 2, 4, 3, 2},
		{1, 1, 3, 2, 4, 3, 2},
		{1, 1, 3, 2, 4, 3, 2},
	}
	threshold1 := 4
	fmt.Println(maxSideLength(mat1, threshold1)) // Output: 2

	mat2 := [][]int{
		{2, 2, 2, 2, 2},
		{2, 2, 2, 2, 2},
		{2, 2, 2, 2, 2},
		{2, 2, 2, 2, 2},
		{2, 2, 2, 2, 2},
	}
	threshold2 := 1
	fmt.Println(maxSideLength(mat2, threshold2)) // Output: 0
}
