package main

import (
	"fmt"
)

func minSwaps(grid [][]int) int {
	n := len(grid)

	// Step 1: Count trailing zeros for each row
	trailingZeros := make([]int, n)
	for i := 0; i < n; i++ {
		count := 0
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] == 0 {
				count++
			} else {
				break
			}
		}
		trailingZeros[i] = count
	}

	swaps := 0

	// Step 2: Fix each row position
	for i := 0; i < n; i++ {
		required := n - i - 1
		j := i

		// Find a row with enough trailing zeros
		for j < n && trailingZeros[j] < required {
			j++
		}

		// If not found, return -1
		if j == n {
			return -1
		}

		// Bring row j up to position i by swapping adjacent rows
		for j > i {
			trailingZeros[j], trailingZeros[j-1] = trailingZeros[j-1], trailingZeros[j]
			swaps++
			j--
		}
	}

	return swaps
}

func main() {
	grid1 := [][]int{
		{0, 0, 1},
		{1, 1, 0},
		{1, 0, 0},
	}
	fmt.Println("Example 1 Output:", minSwaps(grid1)) // Expected: 3

	grid2 := [][]int{
		{0, 1, 1, 0},
		{0, 1, 1, 0},
		{0, 1, 1, 0},
		{0, 1, 1, 0},
	}
	fmt.Println("Example 2 Output:", minSwaps(grid2)) // Expected: -1

	grid3 := [][]int{
		{1, 0, 0},
		{1, 1, 0},
		{1, 1, 1},
	}
	fmt.Println("Example 3 Output:", minSwaps(grid3)) // Expected: 0
}
