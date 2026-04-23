package main

import (
	"fmt"
)

func distance(nums []int) []int64 {
	n := len(nums)
	res := make([]int64, n)

	// Step 1: group indices by value
	groups := make(map[int][]int)
	for i, v := range nums {
		groups[v] = append(groups[v], i)
	}

	// Step 2: process each group
	for _, idx := range groups {
		k := len(idx)
		if k == 1 {
			continue
		}

		// prefix sum of indices
		prefix := make([]int64, k)
		prefix[0] = int64(idx[0])
		for i := 1; i < k; i++ {
			prefix[i] = prefix[i-1] + int64(idx[i])
		}

		for j := 0; j < k; j++ {
			i := idx[j]

			// left contribution
			var left int64
			if j > 0 {
				left = int64(j)*int64(idx[j]) - prefix[j-1]
			}

			// right contribution
			var right int64
			if j < k-1 {
				right = (prefix[k-1] - prefix[j]) - int64(k-j-1)*int64(idx[j])
			}

			res[i] = left + right
		}
	}

	return res
}

func main() {
	// Test cases
	nums1 := []int{1, 3, 1, 1, 2}
	fmt.Println(distance(nums1)) // Expected: [5 0 3 4 0]

	nums2 := []int{0, 5, 3}
	fmt.Println(distance(nums2)) // Expected: [0 0 0]
}
