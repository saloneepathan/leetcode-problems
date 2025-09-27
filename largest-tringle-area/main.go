package main

import (
	"fmt"
	"math"
)

func largestTriangleArea(points [][]int) float64 {
	n := len(points)
	maxArea := 0.0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				x1, y1 := points[i][0], points[i][1]
				x2, y2 := points[j][0], points[j][1]
				x3, y3 := points[k][0], points[k][1]

				area := math.Abs(float64(x1*(y2-y3)+x2*(y3-y1)+x3*(y1-y2))) / 2.0
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}

func main() {
	points1 := [][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}
	fmt.Printf("%.5f\n", largestTriangleArea(points1)) // Expected: 2.00000

	points2 := [][]int{{1, 0}, {0, 0}, {0, 1}}
	fmt.Printf("%.5f\n", largestTriangleArea(points2)) // Expected: 0.50000
}
