package main

import (
	"fmt"
	"sort"
)

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	ans := len(intervals)
	right := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][1] <= right {
			ans--
		} else {
			right = intervals[i][1]
		}
	}

	return ans
}

func main() {
	testCases := [][][]int{
		{{1, 4}, {3, 6}, {2, 8}},
		{{1, 4}, {2, 3}},
		{{1, 2}, {1, 4}, {3, 4}},
		{{0, 10}, {5, 12}},
		{{3, 10}, {4, 10}, {5, 11}},
		{{1, 4}, {2, 5}, {3, 6}},
	}

	for i, intervals := range testCases {
		fmt.Printf("Test Case %d:\n", i+1)
		fmt.Printf("Intervals: %v\n", intervals)
		fmt.Printf("Remaining Intervals: %d\n\n", removeCoveredIntervals(intervals))
	}
}
