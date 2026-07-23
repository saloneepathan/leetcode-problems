package main

import (
	"fmt"
	"math/bits"
)

func uniqueXorTriplets(nums []int) int {
	n := len(nums)

	// If n = 1 or n = 2, the answer is simply n.
	if n < 3 {
		return n
	}

	// For n >= 3, answer = 2^(number of bits required to represent n).
	k := bits.Len(uint(n))
	return 1 << k
}

func main() {
	// Example 1
	nums1 := []int{1, 2}
	fmt.Println(uniqueXorTriplets(nums1)) // Output: 2

	// Example 2
	nums2 := []int{1, 2, 3}
	fmt.Println(uniqueXorTriplets(nums2)) // Output: 4

	// Example 3
	nums3 := []int{1, 2, 3, 4}
	fmt.Println(uniqueXorTriplets(nums3)) // Output: 8

	// Example 4
	nums4 := []int{1}
	fmt.Println(uniqueXorTriplets(nums4)) // Output: 1

	// Example 5
	nums5 := []int{1, 2, 3, 4, 5}
	fmt.Println(uniqueXorTriplets(nums5)) // Output: 8
}
