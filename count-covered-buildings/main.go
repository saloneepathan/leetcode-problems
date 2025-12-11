package main

import "fmt"

func countCoveredBuildings(n int, buildings [][]int) int {
	// rowMin[row] = smallest column in that row
	// rowMax[row] = largest column in that row
	// colMin[col] = smallest row in that column
	// colMax[col] = largest row in that column
	rowMin := make(map[int]int)
	rowMax := make(map[int]int)
	colMin := make(map[int]int)
	colMax := make(map[int]int)

	const INF = int(1e9)

	// initialize mins to +INF and maxs to -INF
	for _, b := range buildings {
		x, y := b[0], b[1]
		if _, ok := rowMin[x]; !ok {
			rowMin[x] = INF
			rowMax[x] = -INF
		}
		if _, ok := colMin[y]; !ok {
			colMin[y] = INF
			colMax[y] = -INF
		}
		if y < rowMin[x] {
			rowMin[x] = y
		}
		if y > rowMax[x] {
			rowMax[x] = y
		}
		if x < colMin[y] {
			colMin[y] = x
		}
		if x > colMax[y] {
			colMax[y] = x
		}
	}

	covered := 0
	for _, b := range buildings {
		x, y := b[0], b[1]
		// any to left:   rowMin[x] < y
		// any to right:  rowMax[x] > y
		// any above:     colMin[y] < x
		// any below:     colMax[y] > x
		if rowMin[x] < y && rowMax[x] > y && colMin[y] < x && colMax[y] > x {
			covered++
		}
	}
	return covered
}

func main() {
	fmt.Println(countCoveredBuildings(3, [][]int{{1, 2}, {2, 2}, {3, 2}, {2, 1}, {2, 3}})) // 1
	fmt.Println(countCoveredBuildings(3, [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}}))         // 0
	fmt.Println(countCoveredBuildings(5, [][]int{{1, 3}, {3, 2}, {3, 3}, {3, 5}, {5, 3}})) // 1
}
