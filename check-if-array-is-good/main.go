package main

import (
	"fmt"
)

func isGood(nums []int) bool {
	n := 0

	// Find maximum element
	for _, num := range nums {
		if num > n {
			n = num
		}
	}

	// Length must be n + 1
	if len(nums) != n+1 {
		return false
	}

	freq := make(map[int]int)

	// Count frequency of each number
	for _, num := range nums {
		freq[num]++
	}

	// Numbers 1 to n-1 must appear exactly once
	for i := 1; i < n; i++ {
		if freq[i] != 1 {
			return false
		}
	}

	// Number n must appear exactly twice
	return freq[n] == 2
}

func main() {
	fmt.Println(isGood([]int{2, 1, 3}))          // false
	fmt.Println(isGood([]int{1, 3, 3, 2}))       // true
	fmt.Println(isGood([]int{1, 1}))             // true
	fmt.Println(isGood([]int{3, 4, 4, 1, 2, 1})) // false
}
