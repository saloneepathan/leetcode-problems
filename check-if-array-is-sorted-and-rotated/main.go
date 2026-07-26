package main

import "fmt"

func check(nums []int) bool {
	count := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		// Compare current element with next element circularly
		if nums[i] > nums[(i+1)%n] {
			count++
		}

		// More than one drop means not sorted & rotated
		if count > 1 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(check([]int{3, 4, 5, 1, 2})) // true
	fmt.Println(check([]int{2, 1, 3, 4}))    // false
	fmt.Println(check([]int{1, 2, 3}))       // true
	fmt.Println(check([]int{1, 1, 1}))       // true
	fmt.Println(check([]int{2, 1}))          // true
}
