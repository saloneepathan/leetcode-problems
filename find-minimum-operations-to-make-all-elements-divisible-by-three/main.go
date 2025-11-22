package main

import (
	"fmt"
)

func minimumOperations(nums []int) int {
	ops := 0
	for _, x := range nums {
		r := x % 3
		if r == 1 || r == 2 {
			ops++
		}
	}
	return ops
}

func main() {
	nums1 := []int{1, 2, 3, 4}
	fmt.Println(minimumOperations(nums1)) // Output: 3

	nums2 := []int{3, 6, 9}
	fmt.Println(minimumOperations(nums2)) // Output: 0
}
