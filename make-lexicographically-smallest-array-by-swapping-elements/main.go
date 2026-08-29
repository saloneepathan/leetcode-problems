package main

import (
	"fmt"
	"sort"
)

func lexicographicallySmallestArray(nums []int, limit int) []int {
	n := len(nums)

	type Pair struct {
		val int
		idx int
	}

	// Store value along with its original index
	arr := make([]Pair, n)

	for i, v := range nums {
		arr[i] = Pair{
			val: v,
			idx: i,
		}
	}

	// Sort by value
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].val < arr[j].val
	})

	start := 0

	for start < n {
		end := start

		// Find the connected component.
		// Consecutive values in sorted order must differ
		// by at most limit.
		for end+1 < n &&
			arr[end+1].val-arr[end].val <= limit {
			end++
		}

		// Collect original indices of this component
		indices := make([]int, end-start+1)

		for i := start; i <= end; i++ {
			indices[i-start] = arr[i].idx
		}

		// Smaller values should go to smaller indices
		sort.Ints(indices)

		// Assign sorted values to sorted indices
		for i := start; i <= end; i++ {
			nums[indices[i-start]] = arr[i].val
		}

		start = end + 1
	}

	return nums
}

func main() {
	// Example 1
	nums1 := []int{1, 5, 3, 9, 8}
	limit1 := 2

	result1 := lexicographicallySmallestArray(nums1, limit1)

	fmt.Println("Example 1:")
	fmt.Println(result1)

	// Example 2
	nums2 := []int{1, 7, 6, 18, 2, 1}
	limit2 := 3

	result2 := lexicographicallySmallestArray(nums2, limit2)

	fmt.Println("Example 2:")
	fmt.Println(result2)

	// Example 3
	nums3 := []int{1, 7, 28, 19, 10}
	limit3 := 3

	result3 := lexicographicallySmallestArray(nums3, limit3)

	fmt.Println("Example 3:")
	fmt.Println(result3)
}