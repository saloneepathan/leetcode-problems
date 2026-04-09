package main

import "fmt"

func xorAfterQueries(nums []int, queries [][]int) int {
	mod := 1000000007

	for _, q := range queries {
		l := q[0]
		r := q[1]
		k := q[2]
		v := q[3]

		for i := l; i <= r; i += k {
			nums[i] = int((int64(nums[i]) * int64(v)) % int64(mod))
		}
	}

	result := 0
	for _, num := range nums {
		result ^= num
	}

	return result
}

func main() {

	// Example 1
	nums1 := []int{1, 1, 1}
	queries1 := [][]int{
		{0, 2, 1, 4},
	}

	fmt.Println("Output:", xorAfterQueries(nums1, queries1)) // 4

	// Example 2
	nums2 := []int{2, 3, 1, 5, 4}
	queries2 := [][]int{
		{1, 4, 2, 3},
		{0, 2, 1, 2},
	}

	fmt.Println("Output:", xorAfterQueries(nums2, queries2)) // 31
}
