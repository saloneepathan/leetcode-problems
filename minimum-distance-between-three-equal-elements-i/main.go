package main

import (
	"fmt"
)

func minimumDistance(nums []int) int {
	indexMap := make(map[int][]int)

	// Store indices of each number
	for i, v := range nums {
		indexMap[v] = append(indexMap[v], i)
	}

	minDist := int(1e18)
	found := false

	// Check each number's indices
	for _, indices := range indexMap {
		if len(indices) < 3 {
			continue
		}

		for i := 0; i+2 < len(indices); i++ {
			dist := 2 * (indices[i+2] - indices[i])
			if dist < minDist {
				minDist = dist
				found = true
			}
		}
	}

	if !found {
		return -1
	}

	return minDist
}

func main() {
	// Test cases
	nums1 := []int{1, 2, 1, 1, 3}
	nums2 := []int{1, 1, 2, 3, 2, 1, 2}
	nums3 := []int{1}

	fmt.Println(minimumDistance(nums1)) // Expected: 6
	fmt.Println(minimumDistance(nums2)) // Expected: 8
	fmt.Println(minimumDistance(nums3)) // Expected: -1
}
