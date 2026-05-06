package main

import (
	"fmt"
)

func rotateTheBox(boxGrid [][]byte) [][]byte {
	m := len(boxGrid)
	n := len(boxGrid[0])

	// Step 1: simulate gravity (stones fall to the right)
	for i := 0; i < m; i++ {
		write := n - 1
		for j := n - 1; j >= 0; j-- {
			if boxGrid[i][j] == '*' {
				write = j - 1
			} else if boxGrid[i][j] == '#' {
				boxGrid[i][j] = '.'
				boxGrid[i][write] = '#'
				write--
			}
		}
	}

	// Step 2: rotate 90° clockwise
	result := make([][]byte, n)
	for i := range result {
		result[i] = make([]byte, m)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			result[j][m-1-i] = boxGrid[i][j]
		}
	}

	return result
}

func printGrid(grid [][]byte) {
	for _, row := range grid {
		for _, ch := range row {
			fmt.Printf("%c ", ch)
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	boxGrid := [][]byte{
		{'#', '.', '*', '.'},
		{'#', '#', '*', '.'},
	}

	fmt.Println("Original:")
	printGrid(boxGrid)

	result := rotateTheBox(boxGrid)

	fmt.Println("Rotated:")
	printGrid(result)
}
