package main

import "fmt"

func rangeAddQueries(n int, queries [][]int) [][]int {
	// diff uses n+1 rows so we can safely apply +1/-1 at r2+1 and c2+1
	diff := make([][]int, n+1)
	for i := range diff {
		diff[i] = make([]int, n+1)
	}

	// Apply difference updates
	for _, q := range queries {
		r1, c1, r2, c2 := q[0], q[1], q[2], q[3]

		diff[r1][c1]++
		diff[r1][c2+1]--
		diff[r2+1][c1]--
		diff[r2+1][c2+1]++
	}

	// Prepare result matrix
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}

	// Build final matrix using 2D prefix sums
	for i := 0; i < n; i++ {
		row := res[i]
		for j := 0; j < n; j++ {
			val := diff[i][j]
			if i > 0 {
				val += res[i-1][j]
			}
			if j > 0 {
				val += row[j-1]
			}
			if i > 0 && j > 0 {
				val -= res[i-1][j-1]
			}
			row[j] = val
		}
	}

	return res
}

func main() {
	n := 3
	queries := [][]int{{1, 1, 2, 2}, {0, 0, 1, 1}}

	ans := rangeAddQueries(n, queries)
	fmt.Println(ans)
}
