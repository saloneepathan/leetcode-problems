package main

import "fmt"

func missingMultiple(nums []int, k int) int {
	seen := make(map[int]bool)

	for _, num := range nums {
		seen[num] = true
	}

	for multiple := k; ; multiple += k {
		if !seen[multiple] {
			return multiple
		}
	}
}

func main() {
	nums1 := []int{8, 2, 3, 4, 6}
	k1 := 2
	fmt.Println(missingMultiple(nums1, k1)) // 10

	nums2 := []int{1, 4, 7, 10, 15}
	k2 := 5
	fmt.Println(missingMultiple(nums2, k2)) // 5
}