package main

import (
	"fmt"
	"sort"
)

func minimumBoxes(apple []int, capacity []int) int {
	// Calculate total apples
	totalApples := 0
	for _, a := range apple {
		totalApples += a
	}

	// Sort capacities in descending order
	sort.Slice(capacity, func(i, j int) bool {
		return capacity[i] > capacity[j]
	})

	// Greedily select boxes
	currentCapacity := 0
	for i, c := range capacity {
		currentCapacity += c
		if currentCapacity >= totalApples {
			return i + 1
		}
	}

	// Guaranteed to be possible
	return len(capacity)
}

func main() {
	// Example 1
	apple1 := []int{1, 3, 2}
	capacity1 := []int{4, 3, 1, 5, 2}
	fmt.Println(minimumBoxes(apple1, capacity1)) // Output: 2

	// Example 2
	apple2 := []int{5, 5, 5}
	capacity2 := []int{2, 4, 2, 7}
	fmt.Println(minimumBoxes(apple2, capacity2)) // Output: 4
}
