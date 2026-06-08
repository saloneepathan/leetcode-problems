package main

import "fmt"

func pivotArray(nums []int, pivot int) []int {
	less := make([]int, 0)
	equal := make([]int, 0)
	greater := make([]int, 0)

	for _, num := range nums {
		if num < pivot {
			less = append(less, num)
		} else if num == pivot {
			equal = append(equal, num)
		} else {
			greater = append(greater, num)
		}
	}

	result := make([]int, 0, len(nums))
	result = append(result, less...)
	result = append(result, equal...)
	result = append(result, greater...)

	return result
}

func main() {
	nums1 := []int{9, 12, 5, 10, 14, 3, 10}
	pivot1 := 10
	fmt.Println(pivotArray(nums1, pivot1)) // [9 5 3 10 10 12 14]

	nums2 := []int{-3, 4, 3, 2}
	pivot2 := 2
	fmt.Println(pivotArray(nums2, pivot2)) // [-3 2 4 3]
}
