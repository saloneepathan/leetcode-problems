package main

import "fmt"

func maxProductPath(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	mod := int64(1e9 + 7)

	maxDp := make([][]int64, m)
	minDp := make([][]int64, m)

	for i := 0; i < m; i++ {
		maxDp[i] = make([]int64, n)
		minDp[i] = make([]int64, n)
	}

	maxDp[0][0] = int64(grid[0][0])
	minDp[0][0] = int64(grid[0][0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if i == 0 && j == 0 {
				continue
			}

			val := int64(grid[i][j])

			maxVal := int64(-1 << 60)
			minVal := int64(1 << 60)

			if i > 0 {
				a := maxDp[i-1][j] * val
				b := minDp[i-1][j] * val

				if a > maxVal {
					maxVal = a
				}
				if b > maxVal {
					maxVal = b
				}

				if a < minVal {
					minVal = a
				}
				if b < minVal {
					minVal = b
				}
			}

			if j > 0 {
				a := maxDp[i][j-1] * val
				b := minDp[i][j-1] * val

				if a > maxVal {
					maxVal = a
				}
				if b > maxVal {
					maxVal = b
				}

				if a < minVal {
					minVal = a
				}
				if b < minVal {
					minVal = b
				}
			}

			maxDp[i][j] = maxVal
			minDp[i][j] = minVal
		}
	}

	result := maxDp[m-1][n-1]

	if result < 0 {
		return -1
	}

	return int(result % mod)
}

func main() {
	grid := [][]int{
		{1, -2, 1},
		{1, -2, 1},
		{3, -4, 1},
	}

	fmt.Println(maxProductPath(grid)) // Output: 8
}
