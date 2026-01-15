package main

import (
	"fmt"
	"sort"
)

func maximizeSquareHoleArea(n int, m int, hBars []int, vBars []int) int {
	maxH := longestConsecutive(hBars) + 1
	maxV := longestConsecutive(vBars) + 1

	side := min(maxH, maxV)
	return side * side
}

func longestConsecutive(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	sort.Ints(arr)

	maxLen := 1
	currLen := 1

	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1]+1 {
			currLen++
		} else {
			currLen = 1
		}
		if currLen > maxLen {
			maxLen = currLen
		}
	}
	return maxLen
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Example 1
	n1, m1 := 2, 1
	hBars1 := []int{2, 3}
	vBars1 := []int{2}
	fmt.Println(maximizeSquareHoleArea(n1, m1, hBars1, vBars1)) // Output: 4

	// Example 2
	n2, m2 := 1, 1
	hBars2 := []int{2}
	vBars2 := []int{2}
	fmt.Println(maximizeSquareHoleArea(n2, m2, hBars2, vBars2)) // Output: 4

	// Example 3
	n3, m3 := 2, 3
	hBars3 := []int{2, 3}
	vBars3 := []int{2, 4}
	fmt.Println(maximizeSquareHoleArea(n3, m3, hBars3, vBars3)) // Output: 4
}
