package main

import (
	"fmt"
)

func getSneakyNumbers(nums []int) []int {
	freq := make(map[int]int)
	result := []int{}

	for _, num := range nums {
		freq[num]++
		if freq[num] == 2 { // Found a duplicate
			result = append(result, num)
			if len(result) == 2 { // Found both sneaky numbers
				break
			}
		}
	}

	return result
}

func main() {
	// Example 1
	nums1 := []int{0, 1, 1, 0}
	fmt.Println(getSneakyNumbers(nums1)) // Output: [0 1]

	// Example 2
	nums2 := []int{0, 3, 2, 1, 3, 2}
	fmt.Println(getSneakyNumbers(nums2)) // Output: [3 2] or [2 3]

	// Example 3
	nums3 := []int{7, 1, 5, 4, 3, 4, 6, 0, 9, 5, 8, 2}
	fmt.Println(getSneakyNumbers(nums3)) // Output: [4 5] or [5 4]
}
