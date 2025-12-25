package main

import (
	"fmt"
	"sort"
)

func maximumHappinessSum(happiness []int, k int) int64 {
	// Sort happiness in descending order
	sort.Slice(happiness, func(i, j int) bool {
		return happiness[i] > happiness[j]
	})

	var result int64 = 0

	for i := 0; i < k; i++ {
		current := happiness[i] - i
		if current <= 0 {
			break
		}
		result += int64(current)
	}

	return result
}

func main() {
	// Example test cases
	fmt.Println(maximumHappinessSum([]int{1, 2, 3}, 2))    // Output: 4
	fmt.Println(maximumHappinessSum([]int{1, 1, 1, 1}, 2)) // Output: 1
	fmt.Println(maximumHappinessSum([]int{2, 3, 4, 5}, 1)) // Output: 5
}
