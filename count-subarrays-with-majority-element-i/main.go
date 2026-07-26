package main

import (
	"fmt"
)

func countMajoritySubarrays(nums []int, target int) int {
	n := len(nums)
	ans := 0

	for left := 0; left < n; left++ {
		targetCount := 0

		for right := left; right < n; right++ {
			if nums[right] == target {
				targetCount++
			}

			length := right - left + 1

			// target is a majority if it appears strictly more than half
			if targetCount*2 > length {
				ans++
			}
		}
	}

	return ans
}

func main() {
	nums1 := []int{1, 2, 2, 3}
	target1 := 2
	fmt.Println(countMajoritySubarrays(nums1, target1)) // 5

	nums2 := []int{1, 1, 1, 1}
	target2 := 1
	fmt.Println(countMajoritySubarrays(nums2, target2)) // 10

	nums3 := []int{1, 2, 3}
	target3 := 4
	fmt.Println(countMajoritySubarrays(nums3, target3)) // 0
}
