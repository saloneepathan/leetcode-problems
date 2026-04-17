package main

import (
	"fmt"
)

func minMirrorPairDistance(nums []int) int {
	indexMap := make(map[int]int, len(nums))
	minDist := len(nums) + 1

	for i, num := range nums {
		// Check if current number matches any stored reversed number
		if prevIdx, ok := indexMap[num]; ok {
			if i-prevIdx < minDist {
				minDist = i - prevIdx
			}
		}

		// Reverse the number
		x := num
		rev := 0
		for x > 0 {
			rev = rev*10 + x%10
			x /= 10
		}

		// Store reversed number with index
		indexMap[rev] = i
	}

	if minDist == len(nums)+1 {
		return -1
	}
	return minDist
}

func main() {
	fmt.Println(minMirrorPairDistance([]int{12, 21, 45, 33, 54})) // Output: 1
	fmt.Println(minMirrorPairDistance([]int{120, 21}))             // Output: 1
	fmt.Println(minMirrorPairDistance([]int{21, 120}))             // Output: -1

	// Additional tests
	fmt.Println(minMirrorPairDistance([]int{11, 11, 11})) // Output: 1
	fmt.Println(minMirrorPairDistance([]int{10, 1}))      // Output: 1
}