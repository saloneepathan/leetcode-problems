package main

import (
	"fmt"
)

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	prefixes := make(map[int]bool)

	// Store all prefixes from arr1
	for _, num := range arr1 {
		for num > 0 {
			prefixes[num] = true
			num /= 10
		}
	}

	ans := 0

	// Check prefixes from arr2
	for _, num := range arr2 {
		for num > 0 {
			if prefixes[num] {
				ans = max(ans, digits(num))
			}
			num /= 10
		}
	}

	return ans
}

// Returns number of digits in num
func digits(num int) int {
	count := 0
	for num > 0 {
		count++
		num /= 10
	}
	return count
}

// Utility max function
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr1 := []int{1, 10, 100}
	arr2 := []int{1000}

	fmt.Println("Output:", longestCommonPrefix(arr1, arr2)) // 3

	arr3 := []int{1, 2, 3}
	arr4 := []int{4, 4, 4}

	fmt.Println("Output:", longestCommonPrefix(arr3, arr4)) // 0
}
