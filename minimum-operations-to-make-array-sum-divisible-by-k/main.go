package main

import (
	"fmt"
)

func minOperations(nums []int, k int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum % k
}

func main() {
	// Example Inputs
	examples := []struct {
		nums []int
		k    int
	}{
		{[]int{3, 9, 7}, 5},
		{[]int{4, 1, 3}, 4},
		{[]int{3, 2}, 6},
	}

	for _, ex := range examples {
		fmt.Printf("nums = %v, k = %d → min operations = %d\n",
			ex.nums, ex.k, minOperations(ex.nums, ex.k))
	}
}
