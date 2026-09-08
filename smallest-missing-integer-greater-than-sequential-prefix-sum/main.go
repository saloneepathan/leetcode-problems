package main

import "fmt"

func missingInteger(nums []int) int {
	// Calculate the sum of the longest sequential prefix.
	sum := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			sum += nums[i]
		} else {
			break
		}
	}

	// Put all numbers in a set for O(1) lookup.
	set := make(map[int]bool)

	for _, num := range nums {
		set[num] = true
	}

	// Find the smallest missing integer >= sum.
	for set[sum] {
		sum++
	}

	return sum
}

func main() {
	nums1 := []int{1, 2, 3, 2, 5}
	fmt.Println(missingInteger(nums1)) // 6

	nums2 := []int{3, 4, 5, 1, 12, 14, 13}
	fmt.Println(missingInteger(nums2)) // 15
}
