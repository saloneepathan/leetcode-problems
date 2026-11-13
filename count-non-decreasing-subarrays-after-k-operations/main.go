package main

import "fmt"

// countNonDecreasingSubarrays counts subarrays that can be made non-decreasing
// with at most k "+1" operations.
//
// Cost of subarray [l..r] = sum over i of (prefixMax(l..i) - nums[i]).
// We sweep l from right to left and keep r as large as the budget allows
// (r only ever moves left). A monotonic stack holds the "records" (prefix-max
// positions) of the window, so each element is pushed/popped O(1) times.
func countNonDecreasingSubarrays(nums []int, k int) int64 {
	n := len(nums)
	stack := make([]int, 0, n) // record positions; stack[head] is the rightmost record
	head := 0
	r := n - 1
	cost := 0
	var ans int64

	for l := n - 1; l >= 0; l-- {
		// Adding nums[l] on the left: every record <= nums[l] gets raised to nums[l].
		for len(stack) > head && nums[stack[len(stack)-1]] <= nums[l] {
			top := len(stack) - 1
			// segment of this record ends right before the next record (or at r)
			end := r
			if top > head {
				end = stack[top-1] - 1
			}
			cost += (nums[l] - nums[stack[top]]) * (end - stack[top] + 1)
			stack = stack[:top]
		}
		stack = append(stack, l)

		// Shrink the right side until the cost fits in k.
		for cost > k {
			b := stack[head] // record owning position r
			cost -= nums[b] - nums[r]
			if b == r {
				head++ // its segment became empty
			}
			r--
		}
		ans += int64(r - l + 1)
	}
	return ans
}

func main() {
	tests := []struct {
		nums []int
		k    int
		want int64
	}{
		{[]int{6, 3, 1, 2, 4, 4}, 7, 17}, // example 1
		{[]int{6, 3, 1, 3, 6}, 4, 12},    // example 2
		{[]int{5}, 1, 1},                 // single element
		{[]int{1, 2, 3, 4}, 1, 10},       // already sorted: all subarrays
		{[]int{4, 3, 2, 1}, 1, 7},        // strictly decreasing, tiny budget
		{[]int{2, 2, 2}, 1, 6},           // all equal
	}
	for _, t := range tests {
		got := countNonDecreasingSubarrays(t.nums, t.k)
		status := "PASS"
		if got != t.want {
			status = "FAIL"
		}
		fmt.Printf("nums=%v k=%d expected=%d actual=%d %s\n", t.nums, t.k, t.want, got, status)
	}
}
