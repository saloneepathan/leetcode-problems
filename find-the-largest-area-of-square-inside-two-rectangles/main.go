package main

import (
	"fmt"
)

func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64 {
	n := len(bottomLeft)
	var maxSide int64 = 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// Intersection boundaries
			xLeft := max(bottomLeft[i][0], bottomLeft[j][0])
			yBottom := max(bottomLeft[i][1], bottomLeft[j][1])
			xRight := min(topRight[i][0], topRight[j][0])
			yTop := min(topRight[i][1], topRight[j][1])

			width := xRight - xLeft
			height := yTop - yBottom

			if width > 0 && height > 0 {
				side := int64(min(width, height))
				if side > maxSide {
					maxSide = side
				}
			}
		}
	}

	return maxSide * maxSide
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example 1
	bottomLeft1 := [][]int{{1, 1}, {2, 2}, {3, 1}}
	topRight1 := [][]int{{3, 3}, {4, 4}, {6, 6}}
	fmt.Println(largestSquareArea(bottomLeft1, topRight1)) // Output: 1

	// Example 2
	bottomLeft2 := [][]int{{1, 1}, {1, 3}, {1, 5}}
	topRight2 := [][]int{{5, 5}, {5, 7}, {5, 9}}
	fmt.Println(largestSquareArea(bottomLeft2, topRight2)) // Output: 4

	// Example 3
	bottomLeft3 := [][]int{{1, 1}, {2, 2}, {1, 2}}
	topRight3 := [][]int{{3, 3}, {4, 4}, {3, 4}}
	fmt.Println(largestSquareArea(bottomLeft3, topRight3)) // Output: 1

	// Example 4
	bottomLeft4 := [][]int{{1, 1}, {3, 3}, {3, 1}}
	topRight4 := [][]int{{2, 2}, {4, 4}, {4, 2}}
	fmt.Println(largestSquareArea(bottomLeft4, topRight4)) // Output: 0
}
