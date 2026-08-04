package main

import "fmt"

func findGCD(nums []int) int {
	minVal, maxVal := nums[0], nums[0]

	// Find minimum and maximum
	for _, num := range nums {
		if num < minVal {
			minVal = num
		}
		if num > maxVal {
			maxVal = num
		}
	}

	// Euclidean Algorithm
	for maxVal%minVal != 0 {
		maxVal, minVal = minVal, maxVal%minVal
	}

	return minVal
}

func main() {
	nums1 := []int{2, 5, 6, 9, 10}
	fmt.Println(findGCD(nums1)) // 2

	nums2 := []int{7, 5, 6, 8, 3}
	fmt.Println(findGCD(nums2)) // 1

	nums3 := []int{3, 3}
	fmt.Println(findGCD(nums3)) // 3
}
