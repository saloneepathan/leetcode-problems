package main

import (
	"fmt"
)

func constructProductMatrix(grid [][]int) [][]int {
	mod := 12345
	n := len(grid)
	m := len(grid[0])
	total := n * m

	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, m)
	}

	// prefix product
	prefix := 1
	for i := 0; i < total; i++ {
		r := i / m
		c := i % m
		res[r][c] = prefix
		prefix = (prefix * (grid[r][c] % mod)) % mod
	}

	// suffix product
	suffix := 1
	for i := total - 1; i >= 0; i-- {
		r := i / m
		c := i % m
		res[r][c] = (res[r][c] * suffix) % mod
		suffix = (suffix * (grid[r][c] % mod)) % mod
	}

	return res
}

func printMatrix(mat [][]int) {
	for _, row := range mat {
		fmt.Println(row)
	}
	fmt.Println()
}

func main() {
	// Testcase 1
	grid1 := [][]int{
		{1, 2},
		{3, 4},
	}
	result1 := constructProductMatrix(grid1)
	fmt.Println("Output 1:")
	printMatrix(result1)

	// Testcase 2
	grid2 := [][]int{
		{12345},
		{2},
		{1},
	}
	result2 := constructProductMatrix(grid2)
	fmt.Println("Output 2:")
	printMatrix(result2)
}
