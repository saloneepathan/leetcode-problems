package main

import "fmt"

func findRotation(mat [][]int, target [][]int) bool {
	n := len(mat)

	for r := 0; r < 4; r++ {
		if isEqual(mat, target) {
			return true
		}
		rotate(mat, n)
	}

	return false
}

func isEqual(mat [][]int, target [][]int) bool {
	n := len(mat)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] != target[i][j] {
				return false
			}
		}
	}
	return true
}

func rotate(mat [][]int, n int) {
	// transpose
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			mat[i][j], mat[j][i] = mat[j][i], mat[i][j]
		}
	}

	// reverse each row
	for i := 0; i < n; i++ {
		left, right := 0, n-1
		for left < right {
			mat[i][left], mat[i][right] = mat[i][right], mat[i][left]
			left++
			right--
		}
	}
}

func main() {
	mat := [][]int{
		{0, 1},
		{1, 0},
	}

	target := [][]int{
		{1, 0},
		{0, 1},
	}

	fmt.Println(findRotation(mat, target)) // true
}
