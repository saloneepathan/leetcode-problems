package main

import (
	"fmt"
)

func minimumDistance(nums []int) int {
	// store last two indices for each number
	last := make(map[int][]int)

	minDist := int(1e18)

	for i, v := range nums {
		last[v] = append(last[v], i)

		if len(last[v]) >= 3 {
			// take last 3 indices
			l := len(last[v])
			first := last[v][l-3]
			third := last[v][l-1]

			dist := 2 * (third - first)
			if dist < minDist {
				minDist = dist
			}
		}
	}

	if minDist == int(1e18) {
		return -1
	}

	return minDist
}

func main() {
	fmt.Println(minimumDistance([]int{1, 2, 1, 1, 3}))       // 6
	fmt.Println(minimumDistance([]int{1, 1, 2, 3, 2, 1, 2})) // 8
	fmt.Println(minimumDistance([]int{1}))                   // -1
}