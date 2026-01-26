package main

import (
	"fmt"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)

	minDiff := int(1e9)
	n := len(arr)

	// Find minimum difference
	for i := 1; i < n; i++ {
		diff := arr[i] - arr[i-1]
		if diff < minDiff {
			minDiff = diff
		}
	}

	// Collect pairs with minimum difference
	result := [][]int{}
	for i := 1; i < n; i++ {
		if arr[i]-arr[i-1] == minDiff {
			result = append(result, []int{arr[i-1], arr[i]})
		}
	}

	return result
}

func main() {
	arr1 := []int{4, 2, 1, 3}
	arr2 := []int{1, 3, 6, 10, 15}
	arr3 := []int{3, 8, -10, 23, 19, -4, -14, 27}

	fmt.Println(minimumAbsDifference(arr1)) // [[1 2] [2 3] [3 4]]
	fmt.Println(minimumAbsDifference(arr2)) // [[1 3]]
	fmt.Println(minimumAbsDifference(arr3)) // [[-14 -10] [19 23] [23 27]]
}
