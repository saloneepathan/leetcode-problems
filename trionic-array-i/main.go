package main

import "fmt"

func isTrionic(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	i := 0

	// 1) strictly increasing
	for i+1 < n && nums[i] < nums[i+1] {
		i++
	}
	// must have at least one increase
	if i == 0 {
		return false
	}

	// 2) strictly decreasing
	j := i
	for j+1 < n && nums[j] > nums[j+1] {
		j++
	}
	// must have at least one decrease
	if j == i {
		return false
	}

	// 3) strictly increasing again
	k := j
	for k+1 < n && nums[k] < nums[k+1] {
		k++
	}
	// must have at least one increase in third phase
	if k == j {
		return false
	}

	// must reach the end
	return k == n-1
}

func main() {
	testCases := [][]int{
		{1, 3, 5, 4, 2, 6}, // true
		{2, 1, 3},          // false
		{3, 7, 1},          // false
		{1, 2, 1, 2},       // true
		{1, 2, 3},          // false
	}

	for _, nums := range testCases {
		fmt.Println(nums, "->", isTrionic(nums))
	}
}
