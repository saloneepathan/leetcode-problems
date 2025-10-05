package main

import "fmt"

func pacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 || len(heights[0]) == 0 {
		return [][]int{}
	}

	m, n := len(heights), len(heights[0])
	pacific := make([][]bool, m)
	atlantic := make([][]bool, m)
	for i := 0; i < m; i++ {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var dfs func(int, int, [][]bool)
	dfs = func(r, c int, visited [][]bool) {
		visited[r][c] = true
		for _, dir := range directions {
			nr, nc := r+dir[0], c+dir[1]
			if nr < 0 || nr >= m || nc < 0 || nc >= n {
				continue
			}
			if visited[nr][nc] {
				continue
			}
			// Flow only from lower/equal to higher
			if heights[nr][nc] < heights[r][c] {
				continue
			}
			dfs(nr, nc, visited)
		}
	}

	// Pacific (top row, left col)
	for c := 0; c < n; c++ {
		dfs(0, c, pacific)
	}
	for r := 0; r < m; r++ {
		dfs(r, 0, pacific)
	}

	// Atlantic (bottom row, right col)
	for c := 0; c < n; c++ {
		dfs(m-1, c, atlantic)
	}
	for r := 0; r < m; r++ {
		dfs(r, n-1, atlantic)
	}

	// Find cells reachable by both
	var result [][]int
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if pacific[r][c] && atlantic[r][c] {
				result = append(result, []int{r, c})
			}
		}
	}

	return result
}

func main() {
	heights := [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}
	fmt.Println(pacificAtlantic(heights))
}
