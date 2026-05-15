package main

import "fmt"

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		// Minimum is in the right half
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			// Minimum is at mid or in the left half
			right = mid
		}
	}

	return nums[left]
}

func main() {
	nums1 := []int{3, 4, 5, 1, 2}
	fmt.Println(findMin(nums1)) // Output: 1

	nums2 := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(findMin(nums2)) // Output: 0

	nums3 := []int{11, 13, 15, 17}
	fmt.Println(findMin(nums3)) // Output: 11
}
