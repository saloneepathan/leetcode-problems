package main

import "fmt"

func leftRightDifference(nums []int) []int {
	n := len(nums)

	total := 0
	for _, v := range nums {
		total += v
	}

	ans := make([]int, n)
	leftSum := 0

	for i, v := range nums {
		rightSum := total - leftSum - v

		diff := leftSum - rightSum
		if diff < 0 {
			diff = -diff
		}

		ans[i] = diff
		leftSum += v
	}

	return ans
}

func main() {
	fmt.Println(leftRightDifference([]int{10, 4, 8, 3})) // [15 1 11 22]
	fmt.Println(leftRightDifference([]int{1}))           // [0]
}
