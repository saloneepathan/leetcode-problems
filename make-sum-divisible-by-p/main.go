package main

import (
	"fmt"
)

func minSubarray(nums []int, p int) int {
	total := 0
	for _, v := range nums {
		total = (total + v) % p
	}

	need := total % p
	if need == 0 {
		return 0
	}

	prefix := 0
	res := len(nums)

	// map from prefix sum mod p → index
	m := map[int]int{}
	m[0] = -1 // prefix before the array begins

	for i, v := range nums {
		prefix = (prefix + v) % p
		target := (prefix - need + p) % p

		if idx, exists := m[target]; exists {
			if i-idx < res {
				res = i - idx
			}
		}

		m[prefix] = i
	}

	if res == len(nums) {
		return -1
	}
	return res
}

func main() {
	fmt.Println(minSubarray([]int{3, 1, 4, 2}, 6)) // Output: 1
	fmt.Println(minSubarray([]int{6, 3, 5, 2}, 9)) // Output: 2
	fmt.Println(minSubarray([]int{1, 2, 3}, 3))    // Output: 0
}
