package main

import "fmt"

func firstStableIndex(nums []int, k int) int {
	n := len(nums)

	if n == 0 {
		return -1
	}

	// minValue[i] = minimum value from nums[i] to nums[n-1]
	minValue := make([]int, n)
	minValue[n-1] = nums[n-1]

	for i := n - 2; i >= 0; i-- {
		if minValue[i+1] < nums[i] {
			minValue[i] = minValue[i+1]
		} else {
			minValue[i] = nums[i]
		}
	}

	maxValue := nums[0]

	for i := 0; i < n; i++ {
		if nums[i] > maxValue {
			maxValue = nums[i]
		}

		if maxValue-minValue[i] <= k {
			return i
		}
	}

	return -1
}

func main() {
	nums := []int{3, 1, 4, 2, 5}
	k := 2

	result := firstStableIndex(nums, k)

	fmt.Println("nums:", nums)
	fmt.Println("k:", k)
	fmt.Println("First stable index:", result)
}