package main

import (
	"fmt"
	"math"
)

// minimumSumSubarray returns the smallest positive sum among all subarrays
// whose length is in [l, r], or -1 if none exists.
// Prefix sums give any subarray sum in O(1): sum(i..j-1) = pre[j] - pre[i].
func minimumSumSubarray(nums []int, l int, r int) int {
	n := len(nums)
	pre := make([]int, n+1)
	for i, v := range nums {
		pre[i+1] = pre[i] + v
	}

	best := math.MaxInt
	for end := l; end <= n; end++ { // end is exclusive
		for length := l; length <= r && length <= end; length++ {
			s := pre[end] - pre[end-length]
			if s > 0 && s < best {
				best = s
			}
		}
	}
	if best == math.MaxInt {
		return -1
	}
	return best
}

func main() {
	tests := []struct {
		nums     []int
		l, r     int
		expected int
	}{
		{[]int{3, -2, 1, 4}, 2, 3, 1},
		{[]int{-2, 2, -3, 1}, 2, 3, -1},
		{[]int{1, 2, 3, 4}, 2, 4, 3},
		{[]int{5}, 1, 1, 5},          // single element, positive
		{[]int{-5}, 1, 1, -1},        // single element, not positive
		{[]int{0, 0, 0}, 1, 3, -1},   // zeros are not > 0
		{[]int{2, -2, 1}, 1, 3, 1},   // single-element window
		{[]int{1000, -999}, 2, 2, 1}, // extremes
	}
	for _, t := range tests {
		got := minimumSumSubarray(t.nums, t.l, t.r)
		status := "PASS"
		if got != t.expected {
			status = "FAIL"
		}
		fmt.Printf("nums=%v l=%d r=%d expected=%d actual=%d %s\n", t.nums, t.l, t.r, t.expected, got, status)
	}
}
