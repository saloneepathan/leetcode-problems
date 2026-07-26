package main

import "fmt"

func maximumProduct(nums []int) int {
	max1, max2, max3 := -1001, -1001, -1001
	min1, min2 := 1001, 1001

	for _, x := range nums {
		// Update the three largest numbers.
		if x > max1 {
			max3 = max2
			max2 = max1
			max1 = x
		} else if x > max2 {
			max3 = max2
			max2 = x
		} else if x > max3 {
			max3 = x
		}

		// Update the two smallest numbers.
		if x < min1 {
			min2 = min1
			min1 = x
		} else if x < min2 {
			min2 = x
		}
	}

	prod1 := max1 * max2 * max3
	prod2 := max1 * min1 * min2

	if prod1 > prod2 {
		return prod1
	}
	return prod2
}

func main() {
	testCases := [][]int{
		{1, 2, 3},
		{1, 2, 3, 4},
		{-1, -2, -3},
		{-100, -98, -1, 2, 3, 4},
		{-10, -10, 5, 2},
	}

	for _, nums := range testCases {
		fmt.Printf("nums = %v -> Maximum Product = %d\n", nums, maximumProduct(nums))
	}
}
