package main

import (
	"fmt"
	"sort"
)

func countMajoritySubarrays(nums []int, target int) int64 {
	n := len(nums)

	// Prefix balance:
	// +1 if nums[i] == target
	// -1 otherwise
	//
	// For a subarray:
	// target is majority iff balance > 0.
	prefix := 0

	// Coordinate compression of prefix values.
	values := make([]int, 0, n+1)
	values = append(values, 0)

	cur := 0
	for _, x := range nums {
		if x == target {
			cur++
		} else {
			cur--
		}
		values = append(values, cur)
	}

	sort.Ints(values)
	k := 0
	for _, v := range values {
		if k == 0 || values[k-1] != v {
			values[k] = v
			k++
		}
	}
	values = values[:k]

	getIndex := func(x int) int {
		return sort.SearchInts(values, x) + 1 // 1-based index
	}

	// Fenwick Tree
	bit := make([]int, k+2)

	add := func(idx, delta int) {
		for idx <= k {
			bit[idx] += delta
			idx += idx & -idx
		}
	}

	query := func(idx int) int {
		res := 0
		for idx > 0 {
			res += bit[idx]
			idx -= idx & -idx
		}
		return res
	}

	add(getIndex(0), 1)

	var ans int64

	for _, x := range nums {
		if x == target {
			prefix++
		} else {
			prefix--
		}

		// Count previous prefix sums strictly smaller.
		pos := sort.SearchInts(values, prefix)
		ans += int64(query(pos))

		add(getIndex(prefix), 1)
	}

	return ans
}

func main() {
	fmt.Println(countMajoritySubarrays([]int{1, 2, 2, 3}, 2)) // 5
	fmt.Println(countMajoritySubarrays([]int{1, 1, 1, 1}, 1)) // 10
	fmt.Println(countMajoritySubarrays([]int{1, 2, 3}, 4))    // 0
}
