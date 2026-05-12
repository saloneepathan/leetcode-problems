package main

import (
	"fmt"
	"sort"
)

func minimumEffort(tasks [][]int) int {
	// Sort tasks by (minimum - actual) in descending order
	sort.Slice(tasks, func(i, j int) bool {
		return (tasks[i][1] - tasks[i][0]) > (tasks[j][1] - tasks[j][0])
	})

	initialEnergy := 0
	currentEnergy := 0

	for _, task := range tasks {
		actual := task[0]
		minimum := task[1]

		// If current energy is less than required minimum,
		// add the deficit to initial energy
		if currentEnergy < minimum {
			initialEnergy += (minimum - currentEnergy)
			currentEnergy = minimum
		}

		// Perform the task
		currentEnergy -= actual
	}

	return initialEnergy
}

func main() {
	tasks1 := [][]int{{1, 2}, {2, 4}, {4, 8}}
	fmt.Println(minimumEffort(tasks1)) // Output: 8

	tasks2 := [][]int{{1, 3}, {2, 4}, {10, 11}, {10, 12}, {8, 9}}
	fmt.Println(minimumEffort(tasks2)) // Output: 32

	tasks3 := [][]int{{1, 7}, {2, 8}, {3, 9}, {4, 10}, {5, 11}, {6, 12}}
	fmt.Println(minimumEffort(tasks3)) // Output: 27
}
