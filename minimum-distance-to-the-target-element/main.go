package main

import (
	"fmt"
)

func getMinDistance(nums []int, target int, start int) int {
	minDist := 1<<31 - 1 // max int

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			dist := abs(i - start)
			if dist < minDist {
				minDist = dist
			}
		}
	}

	return minDist
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Testcase 1
	nums1 := []int{1, 2, 3, 4, 5}
	target1 := 5
	start1 := 3
	fmt.Println(getMinDistance(nums1, target1, start1)) // Output: 1

	// Testcase 2
	nums2 := []int{1}
	target2 := 1
	start2 := 0
	fmt.Println(getMinDistance(nums2, target2, start2)) // Output: 0

	// Testcase 3
	nums3 := []int{1, 1, 1, 1, 1}
	target3 := 1
	start3 := 0
	fmt.Println(getMinDistance(nums3, target3, start3)) // Output: 0
}
