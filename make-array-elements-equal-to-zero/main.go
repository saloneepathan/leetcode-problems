package main

import "fmt"

func countValidSelections(nums []int) int {
	n := len(nums)
	valid := 0

	// simulate the process
	simulate := func(start int, dir int) bool {
		arr := make([]int, n)
		copy(arr, nums)
		curr := start

		for curr >= 0 && curr < n {
			if arr[curr] == 0 {
				curr += dir
			} else {
				arr[curr]--
				dir = -dir
				curr += dir
			}
		}

		// check if all elements became zero
		for _, v := range arr {
			if v != 0 {
				return false
			}
		}
		return true
	}

	for i, v := range nums {
		if v == 0 {
			if simulate(i, 1) {
				valid++
			}
			if simulate(i, -1) {
				valid++
			}
		}
	}

	return valid
}

func main() {
	fmt.Println(countValidSelections([]int{1, 0, 2, 0, 3}))       // Output: 2
	fmt.Println(countValidSelections([]int{2, 3, 4, 0, 4, 1, 0})) // Output: 0
}
