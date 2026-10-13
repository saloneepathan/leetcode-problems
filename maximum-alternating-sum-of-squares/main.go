package main

import (
	"fmt"
	"sort"
)

func maxAlternatingSum(nums []int) int64 {
	n := len(nums)

	for i := range nums {
		nums[i] *= nums[i]
	}

	sort.Ints(nums)

	var ans int64

	for i := 0; i < n/2; i++ {
		ans -= int64(nums[i])
	}

	for i := n / 2; i < n; i++ {
		ans += int64(nums[i])
	}

	return ans
}

func main() {
	nums1 := []int{1, 2, 3}
	fmt.Println(maxAlternatingSum(nums1)) // Output: 12

	nums2 := []int{1, -1, 2, -2, 3, -3}
	fmt.Println(maxAlternatingSum(nums2)) // Output: 16
}