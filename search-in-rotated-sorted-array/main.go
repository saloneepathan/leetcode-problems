package main

import "fmt"

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		// If target is found
		if nums[mid] == target {
			return mid
		}

		// Check if left half is sorted
		if nums[left] <= nums[mid] {

			// Target lies in left half
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}

		} else {
			// Right half is sorted

			// Target lies in right half
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

func main() {
	nums1 := []int{4, 5, 6, 7, 0, 1, 2}
	target1 := 0
	fmt.Println("Output:", search(nums1, target1)) // 4

	nums2 := []int{4, 5, 6, 7, 0, 1, 2}
	target2 := 3
	fmt.Println("Output:", search(nums2, target2)) // -1

	nums3 := []int{1}
	target3 := 0
	fmt.Println("Output:", search(nums3, target3)) // -1
}
