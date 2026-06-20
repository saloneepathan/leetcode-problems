package main

import (
	"fmt"
	"sort"
)

func maxBuilding(n int, restrictions [][]int) int {
	// Building 1 must always have height 0.
	restrictions = append(restrictions, []int{1, 0})

	// Add building n as a final boundary if it is not already restricted.
	hasN := false
	for _, r := range restrictions {
		if r[0] == n {
			hasN = true
			break
		}
	}

	if !hasN {
		restrictions = append(restrictions, []int{n, n - 1})
	}

	// Sort restrictions by building number.
	sort.Slice(restrictions, func(i, j int) bool {
		return restrictions[i][0] < restrictions[j][0]
	})

	// Left to right: height cannot increase by more than distance.
	for i := 1; i < len(restrictions); i++ {
		distance := restrictions[i][0] - restrictions[i-1][0]
		restrictions[i][1] = min(
			restrictions[i][1],
			restrictions[i-1][1]+distance,
		)
	}

	// Right to left: height cannot increase by more than distance.
	for i := len(restrictions) - 2; i >= 0; i-- {
		distance := restrictions[i+1][0] - restrictions[i][0]
		restrictions[i][1] = min(
			restrictions[i][1],
			restrictions[i+1][1]+distance,
		)
	}

	maxHeight := 0

	// Calculate maximum possible peak between every consecutive restriction pair.
	for i := 1; i < len(restrictions); i++ {
		leftBuilding := restrictions[i-1][0]
		leftHeight := restrictions[i-1][1]

		rightBuilding := restrictions[i][0]
		rightHeight := restrictions[i][1]

		distance := rightBuilding - leftBuilding

		// Peak occurs where the upward and downward slopes meet.
		peak := (leftHeight + rightHeight + distance) / 2

		if peak > maxHeight {
			maxHeight = peak
		}
	}

	return maxHeight
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Example 1
	n1 := 5
	restrictions1 := [][]int{
		{2, 1},
		{4, 1},
	}
	fmt.Println(maxBuilding(n1, restrictions1)) // Output: 2

	// Example 2
	n2 := 6
	restrictions2 := [][]int{}
	fmt.Println(maxBuilding(n2, restrictions2)) // Output: 5

	// Example 3
	n3 := 10
	restrictions3 := [][]int{
		{5, 3},
		{2, 5},
		{7, 4},
		{10, 3},
	}
	fmt.Println(maxBuilding(n3, restrictions3)) // Output: 5
}
