package main

import (
	"fmt"
)

func maxRunTime(n int, batteries []int) int64 {
	var total int64 = 0
	for _, b := range batteries {
		total += int64(b)
	}

	// Binary search for the maximum time
	left, right := int64(0), total/int64(n)

	for left < right {
		mid := (left + right + 1) / 2
		if canRun(mid, n, batteries) {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return left
}

func canRun(t int64, n int, batteries []int) bool {
	var usable int64 = 0
	for _, b := range batteries {
		if int64(b) > t {
			usable += t
		} else {
			usable += int64(b)
		}
	}
	return usable >= t*int64(n)
}

func main() {
	n1 := 2
	b1 := []int{3, 3, 3}
	fmt.Println(maxRunTime(n1, b1)) // Expected: 4

	n2 := 2
	b2 := []int{1, 1, 1, 1}
	fmt.Println(maxRunTime(n2, b2)) // Expected: 2
}
