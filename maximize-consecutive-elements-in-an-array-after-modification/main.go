package main

import (
	"fmt"
	"sort"
)

func maxSelectedElements(nums []int) int {
	sort.Ints(nums)

	dp := map[int]int{}
	ans := 0

	for _, x := range nums {
		dp[x+1] = max(dp[x+1], dp[x]+1)
		dp[x] = max(dp[x], dp[x-1]+1)

		ans = max(ans, dp[x], dp[x+1])
	}

	return ans
}

func max(values ...int) int {
	result := values[0]

	for _, value := range values[1:] {
		if value > result {
			result = value
		}
	}

	return result
}

func main() {
	nums1 := []int{2, 1, 5, 1, 1}
	fmt.Println(maxSelectedElements(nums1)) // 3

	nums2 := []int{1, 4, 7, 10}
	fmt.Println(maxSelectedElements(nums2)) // 1

	nums3 := []int{
		8, 10, 6, 12, 9, 12, 2,
		3, 13, 19, 11, 18, 10, 16,
	}
	fmt.Println(maxSelectedElements(nums3)) // 8
}
