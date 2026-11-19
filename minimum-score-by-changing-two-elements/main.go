package main

import (
	"fmt"
	"sort"
)

func minimizeSum(nums []int) int {
	sort.Ints(nums)

	n := len(nums)

	// Change:
	// 1. Two smallest
	// 2. Smallest + largest
	// 3. Largest + second largest
	// 4. Second smallest + second largest
	//
	// After changing two elements, we can make the
	// minimum difference 0, so we only need to
	// minimize the maximum difference.

	ans := nums[n-1] - nums[2]

	ans = min(ans, nums[n-2]-nums[1])
	ans = min(ans, nums[n-3]-nums[0])
	ans = min(ans, nums[n-2]-nums[1])

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 4, 7, 8, 5}

	fmt.Println(minimizeSum(nums))
}