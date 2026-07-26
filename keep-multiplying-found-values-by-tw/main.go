package main

import (
	"fmt"
)

func findFinalValue(nums []int, original int) int {
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}

	for set[original] {
		original *= 2
	}

	return original
}

func main() {
	nums := []int{5, 3, 6, 1, 12}
	original := 3

	fmt.Println(findFinalValue(nums, original)) // Output: 24
}
