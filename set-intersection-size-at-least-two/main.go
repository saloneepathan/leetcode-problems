package main

import (
	"fmt"
	"sort"
)

func intersectionSizeTwo(intervals [][]int) int {
	// Sort by end ascending, and start descending when ends tie
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] == intervals[j][1] {
			return intervals[i][0] > intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})

	// last two chosen points
	a, b := -1, -1
	ans := 0

	for _, in := range intervals {
		start, end := in[0], in[1]

		if b < start {
			// no overlap → need two new points
			ans += 2
			a = end - 1
			b = end
		} else if a < start {
			// one point overlaps → need one more point
			ans += 1
			a = b
			b = end
		}
		// else: already covered by a and b
	}

	return ans
}

func main() {
	intervals1 := [][]int{{1, 3}, {3, 7}, {8, 9}}
	fmt.Println(intersectionSizeTwo(intervals1)) // Output: 5

	intervals2 := [][]int{{1, 3}, {1, 4}, {2, 5}, {3, 5}}
	fmt.Println(intersectionSizeTwo(intervals2)) // Output: 3

	intervals3 := [][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}}
	fmt.Println(intersectionSizeTwo(intervals3)) // Output: 5
}
