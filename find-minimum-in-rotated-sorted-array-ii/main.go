package main

import "fmt"

func findMin(nums []int) int {
	n := len(nums) - 1
	last := nums[n]

	left, right := 0, n

	for left < n && nums[left] == last {
		left++
	}

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > last {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}

	ans := findMin(nums)

	fmt.Println(ans) // Output: 0
}
