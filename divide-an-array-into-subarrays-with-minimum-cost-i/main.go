package main

import (
	"fmt"
)

func minimumCost(nums []int) int {
	first := nums[0]

	// Find the two smallest values in nums[1:]
	min1, min2 := 101, 101 // nums[i] <= 50

	for i := 1; i < len(nums); i++ {
		if nums[i] < min1 {
			min2 = min1
			min1 = nums[i]
		} else if nums[i] < min2 {
			min2 = nums[i]
		}
	}

	return first + min1 + min2
}

func main() {
	fmt.Println(minimumCost([]int{1, 2, 3, 12})) // 6
	fmt.Println(minimumCost([]int{5, 4, 3}))     // 12
	fmt.Println(minimumCost([]int{10, 3, 1, 1})) // 12
}
