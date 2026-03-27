package main

import (
	"fmt"
)

func areSimilar(mat [][]int, k int) bool {
	m := len(mat)
	n := len(mat[0])

	k = k % n

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i%2 == 0 {
				// even row -> left shift
				if mat[i][j] != mat[i][(j+k)%n] {
					return false
				}
			} else {
				// odd row -> right shift
				if mat[i][j] != mat[i][(j-k+n)%n] {
					return false
				}
			}
		}
	}

	return true
}

func main() {

	// Example 1
	mat1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	k1 := 4
	fmt.Println(areSimilar(mat1, k1)) // false

	// Example 2
	mat2 := [][]int{
		{1, 2, 1, 2},
		{5, 5, 5, 5},
		{6, 3, 6, 3},
	}
	k2 := 2
	fmt.Println(areSimilar(mat2, k2)) // true

	// Example 3
	mat3 := [][]int{
		{2, 2},
		{2, 2},
	}
	k3 := 3
	fmt.Println(areSimilar(mat3, k3)) // true
}
