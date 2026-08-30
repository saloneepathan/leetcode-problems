package main

import (
	"bufio"
	"fmt"
	"os"
)

func minimumDeletions(nums []int) int {
	n := len(nums)

	minIdx, maxIdx := 0, 0

	// Find indices of minimum and maximum elements.
	for i := 1; i < n; i++ {
		if nums[i] < nums[minIdx] {
			minIdx = i
		}
		if nums[i] > nums[maxIdx] {
			maxIdx = i
		}
	}

	// Make left the smaller index and right the larger index.
	left, right := minIdx, maxIdx
	if left > right {
		left, right = right, left
	}

	// Option 1: Remove both from the front.
	fromFront := right + 1

	// Option 2: Remove both from the back.
	fromBack := n - left

	// Option 3: Remove one from the front and one from the back.
	fromBothEnds := (left + 1) + (n - right)

	return min(fromFront, min(fromBack, fromBothEnds))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &nums[i])
	}

	result := minimumDeletions(nums)

	fmt.Fprintln(out, result)
}