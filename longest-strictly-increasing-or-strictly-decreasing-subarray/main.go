package main

import (
	"fmt"
	"math/rand"
)

// longestMonotonicSubarray tracks the length of the current strictly
// increasing run and the current strictly decreasing run in one pass.
func longestMonotonicSubarray(nums []int) int {
	best, inc, dec := 1, 1, 1
	for i := 1; i < len(nums); i++ {
		switch {
		case nums[i] > nums[i-1]:
			inc++
			dec = 1
		case nums[i] < nums[i-1]:
			dec++
			inc = 1
		default: // equal breaks both runs
			inc, dec = 1, 1
		}
		best = max(best, inc, dec)
	}
	return best
}

// brute force: check every subarray (used only for cross-checking).
func brute(nums []int) int {
	best := 1
	for i := range nums {
		for j := i; j < len(nums); j++ {
			incOK, decOK := true, true
			for k := i + 1; k <= j; k++ {
				if nums[k] <= nums[k-1] {
					incOK = false
				}
				if nums[k] >= nums[k-1] {
					decOK = false
				}
			}
			if (incOK || decOK) && j-i+1 > best {
				best = j - i + 1
			}
		}
	}
	return best
}

func main() {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 4, 3, 3, 2}, 2},
		{[]int{3, 3, 3, 3}, 1},
		{[]int{3, 2, 1}, 3},
		{[]int{7}, 1},                // single element
		{[]int{1, 2, 3, 4, 5}, 5},    // all increasing
		{[]int{1, 2, 3, 2, 1, 0}, 4}, // peak: decreasing run 3,2,1,0
		{[]int{5, 4, 4, 5, 6, 7}, 4}, // equal breaks, then increasing
	}
	for _, t := range tests {
		got := longestMonotonicSubarray(t.nums)
		status := "PASS"
		if got != t.expected {
			status = "FAIL"
		}
		fmt.Printf("input=%v expected=%d actual=%d %s\n", t.nums, t.expected, got, status)
	}

	// random cross-check against brute force
	r := rand.New(rand.NewSource(1))
	bad := 0
	for i := 0; i < 2000; i++ {
		n := r.Intn(12) + 1
		a := make([]int, n)
		for j := range a {
			a[j] = r.Intn(6) + 1
		}
		if longestMonotonicSubarray(a) != brute(a) {
			bad++
			fmt.Println("MISMATCH", a)
		}
	}
	fmt.Printf("random cross-check: 2000 cases, mismatches=%d\n", bad)
}
