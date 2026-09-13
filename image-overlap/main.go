package main

import "fmt"

func largestOverlap(img1 [][]int, img2 [][]int) int {
	n := len(img1)
	maxOverlap := 0

	// Try every possible translation (dr, dc).
	for dr := -(n - 1); dr <= n-1; dr++ {
		for dc := -(n - 1); dc <= n-1; dc++ {
			overlap := 0

			for r := 0; r < n; r++ {
				for c := 0; c < n; c++ {
					r2 := r + dr
					c2 := c + dc

					// Check if translated position is inside img2.
					if r2 >= 0 && r2 < n && c2 >= 0 && c2 < n {
						if img1[r][c] == 1 && img2[r2][c2] == 1 {
							overlap++
						}
					}
				}
			}

			if overlap > maxOverlap {
				maxOverlap = overlap
			}
		}
	}

	return maxOverlap
}

func main() {
	img1 := [][]int{
		{1, 1, 0},
		{0, 1, 0},
		{0, 1, 0},
	}

	img2 := [][]int{
		{0, 0, 0},
		{0, 1, 1},
		{0, 0, 1},
	}

	fmt.Println(largestOverlap(img1, img2)) // 3
}