package main

import (
	"fmt"
	"sort"
)

func getBiggestThree(grid [][]int) []int {
	m, n := len(grid), len(grid[0])

	// diagonal prefix sums
	d1 := make([][]int, m+1)
	d2 := make([][]int, m+1)

	for i := range d1 {
		d1[i] = make([]int, n+2)
		d2[i] = make([]int, n+2)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			d1[i][j] = grid[i-1][j-1] + d1[i-1][j-1]
		}
	}

	for i := 1; i <= m; i++ {
		for j := n; j >= 1; j-- {
			d2[i][j] = grid[i-1][j-1] + d2[i-1][j+1]
		}
	}

	set := map[int]bool{}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {

			set[grid[r][c]] = true // size 0 rhombus

			for k := 1; ; k++ {
				if r-k < 0 || r+k >= m || c-k < 0 || c+k >= n {
					break
				}

				top := r - k + 1
				bottom := r + k + 1
				left := c - k + 1
				right := c + k + 1
				cr := r + 1
				cc := c + 1

				sum := 0

				// top -> right
				sum += d1[cr][right] - d1[top-1][cc-1]

				// right -> bottom
				sum += d2[bottom][cc] - d2[cr-1][right+1]

				// bottom -> left
				sum += d1[bottom][cc] - d1[cr-1][left-1]

				// left -> top
				sum += d2[cr][left] - d2[top-1][cc+1]

				// remove double counted corners
				sum -= grid[r-k][c]
				sum -= grid[r][c+k]
				sum -= grid[r+k][c]
				sum -= grid[r][c-k]

				set[sum] = true
			}
		}
	}

	res := []int{}
	for v := range set {
		res = append(res, v)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(res)))

	if len(res) > 3 {
		res = res[:3]
	}

	return res
}

func main() {
	grid := [][]int{
		{3, 4, 5, 1, 3},
		{3, 3, 4, 2, 3},
		{20, 30, 200, 40, 10},
		{1, 5, 5, 4, 1},
		{4, 3, 2, 2, 5},
	}

	fmt.Println(getBiggestThree(grid))
}
