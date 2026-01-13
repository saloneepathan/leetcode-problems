package main

import (
	"fmt"
)

func separateSquares(squares [][]int) float64 {
	var totalArea float64
	low := float64(1e18)
	high := float64(0)

	// Compute total area and search bounds
	for _, s := range squares {
		y := float64(s[1])
		l := float64(s[2])

		totalArea += l * l

		if y < low {
			low = y
		}
		if y+l > high {
			high = y + l
		}
	}

	target := totalArea / 2.0

	// Binary search for the minimum y
	for i := 0; i < 70; i++ { // 70 iterations ensures precision < 1e-6
		mid := (low + high) / 2.0
		var below float64

		for _, s := range squares {
			y := float64(s[1])
			l := float64(s[2])

			if mid <= y {
				// Entire square is above the line
				continue
			} else if mid >= y+l {
				// Entire square is below the line
				below += l * l
			} else {
				// Line cuts the square
				below += (mid - y) * l
			}
		}

		if below >= target {
			high = mid
		} else {
			low = mid
		}
	}

	return high
}

func main() {
	// Example 1
	squares1 := [][]int{{0, 0, 1}, {2, 2, 1}}
	fmt.Printf("%.5f\n", separateSquares(squares1)) // Expected: 1.00000

	// Example 2
	squares2 := [][]int{{0, 0, 2}, {1, 1, 1}}
	fmt.Printf("%.5f\n", separateSquares(squares2)) // Expected: 1.16667
}
