package main

import "fmt"

// minCost returns the minimum total time required
// to make the rope colorful (no two adjacent balloons
// of the same color).
func minCost(colors string, neededTime []int) int {
	total := 0
	n := len(colors)

	prevColor := colors[0]
	prevTime := neededTime[0]

	for i := 1; i < n; i++ {
		if colors[i] == prevColor {
			// Same color → remove the one with smaller neededTime
			if neededTime[i] < prevTime {
				total += neededTime[i] // remove current
			} else {
				total += prevTime        // remove previous
				prevTime = neededTime[i] // keep current (more costly)
			}
		} else {
			// Move to next color group
			prevColor = colors[i]
			prevTime = neededTime[i]
		}
	}

	return total
}

func main() {
	// Test cases
	fmt.Println(minCost("abaac", []int{1, 2, 3, 4, 5}))                          // Expected: 3
	fmt.Println(minCost("abc", []int{1, 2, 3}))                                  // Expected: 0
	fmt.Println(minCost("aabaa", []int{1, 2, 3, 4, 1}))                          // Expected: 2
	fmt.Println(minCost("aaabbbabbbb", []int{3, 5, 10, 7, 5, 3, 5, 5, 4, 8, 1})) // Expected: 26
	fmt.Println(minCost("aaaa", []int{1, 2, 3, 4}))                              // Expected: 6
}
