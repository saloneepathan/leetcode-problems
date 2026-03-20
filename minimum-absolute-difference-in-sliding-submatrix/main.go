package main

import (
	"fmt"
	"math"
	"sort"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minAbsDiff(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	res := make([][]int, m-k+1)

	for i := range res {
		res[i] = make([]int, n-k+1)
	}

	for i := 0; i+k <= m; i++ {
		for j := 0; j+k <= n; j++ {

			kgrid := []int{}

			// collect elements
			for x := i; x < i+k; x++ {
				for y := j; y < j+k; y++ {
					kgrid = append(kgrid, grid[x][y])
				}
			}

			sort.Ints(kgrid)

			kmin := math.MaxInt

			// compute min absolute difference
			for t := 1; t < len(kgrid); t++ {
				if kgrid[t] == kgrid[t-1] {
					kmin = 0
					break
				}
				kmin = min(kmin, kgrid[t]-kgrid[t-1])
			}

			// if all same OR single element → 0
			if kmin == math.MaxInt {
				res[i][j] = 0
			} else {
				res[i][j] = kmin
			}
		}
	}

	return res
}

func main() {
	grid := [][]int{
		{1, -2, 3},
		{2, 3, 5},
	}
	k := 2

	result := minAbsDiff(grid, k)

	for _, row := range result {
		fmt.Println(row)
	}
}
