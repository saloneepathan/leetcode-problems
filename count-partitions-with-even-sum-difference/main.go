package main

import (
	"fmt"
)

func countPartitions(nums []int) int {
	n := len(nums)
	total := 0
	for _, v := range nums {
		total += v
	}

	// If total sum is even, all (n-1) partitions are valid.
	if total%2 == 0 {
		return n - 1
	}

	// If total is odd, no partition gives even difference.
	return 0
}

func main() {
	// Test cases
	fmt.Println(countPartitions([]int{10, 10, 3, 7, 6})) // Expected 4
	fmt.Println(countPartitions([]int{1, 2, 2}))         // Expected 0
	fmt.Println(countPartitions([]int{2, 4, 6, 8}))      // Expected 3
}
