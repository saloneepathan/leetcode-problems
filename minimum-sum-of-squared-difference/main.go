
package main

import (
	"fmt"
	"sort"
)

func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
	n := len(nums1)
	diff := make([]int, n)

	for i := 0; i < n; i++ {
		diff[i] = abs(nums1[i] - nums2[i])
	}

	k := k1 + k2

	sort.Sort(sort.Reverse(sort.IntSlice(diff)))

	for k > 0 && diff[0] > 0 {
		diff[0]--
		k--

		// Keep the differences sorted in descending order.
		for i := 0; i < n-1; i++ {
			if diff[i] < diff[i+1] {
				diff[i], diff[i+1] = diff[i+1], diff[i]
			} else {
				break
			}
		}
	}

	var ans int64
	for _, d := range diff {
		ans += int64(d) * int64(d)
	}

	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	nums1 := []int{1, 4, 10, 12}
	nums2 := []int{5, 8, 6, 9}

	fmt.Println(minSumSquareDiff(nums1, nums2, 1, 1))
}
