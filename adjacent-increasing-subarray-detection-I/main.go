package main

import "fmt"

func hasIncreasingSubarrays(nums []int, k int) bool {
	n := len(nums)
	if 2*k > n {
		return false
	}

	isIncreasing := func(start int) bool {
		for i := start; i < start+k-1; i++ {
			if nums[i] >= nums[i+1] {
				return false
			}
		}
		return true
	}

	for i := 0; i+k <= n-k; i++ {
		if isIncreasing(i) && isIncreasing(i+k) {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(hasIncreasingSubarrays([]int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1}, 3)) // true
	fmt.Println(hasIncreasingSubarrays([]int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7}, 5)) // false
}
