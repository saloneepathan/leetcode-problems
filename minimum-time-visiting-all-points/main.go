package main

import "fmt"

func minTimeToVisitAllPoints(points [][]int) int {
	totalTime := 0

	for i := 1; i < len(points); i++ {
		dx := abs(points[i][0] - points[i-1][0])
		dy := abs(points[i][1] - points[i-1][1])
		totalTime += max(dx, dy)
	}

	return totalTime
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	points1 := [][]int{{1, 1}, {3, 4}, {-1, 0}}
	fmt.Println(minTimeToVisitAllPoints(points1)) // Output: 7

	points2 := [][]int{{3, 2}, {-2, 2}}
	fmt.Println(minTimeToVisitAllPoints(points2)) // Output: 5
}
