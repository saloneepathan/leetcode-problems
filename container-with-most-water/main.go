package main

import (
	"fmt"
)

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxWater := 0

	for left < right {
		// Calculate the width and height of the current container
		width := right - left
		h := min(height[left], height[right])

		// Calculate current area
		area := width * h
		if area > maxWater {
			maxWater = area
		}

		// Move the pointer with the smaller height inward
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxWater
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height)) // Output: 49
}
