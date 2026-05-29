package main

import "fmt"

func minElement(nums []int) int {
	minVal := 1<<31 - 1

	for _, n := range nums {
		sum := 0

		for n > 0 {
			sum += n % 10
			n /= 10
		}

		if sum < minVal {
			minVal = sum
		}
	}

	return minVal
}

func main() {
	nums1 := []int{10, 12, 13, 14}
	fmt.Println(minElement(nums1)) // Output: 1

	nums2 := []int{1, 2, 3, 4}
	fmt.Println(minElement(nums2)) // Output: 1

	nums3 := []int{999, 19, 199}
	fmt.Println(minElement(nums3)) // Output: 10
}
