package main

import (
	"fmt"
)

func prefixesDivBy5(nums []int) []bool {
	ans := make([]bool, len(nums))
	curr := 0

	for i, bit := range nums {
		curr = (curr*2 + bit) % 5
		ans[i] = (curr == 0)
	}

	return ans
}

func main() {
	nums := []int{0, 1, 1}
	result := prefixesDivBy5(nums)

	fmt.Println("Input:", nums)
	fmt.Println("Output:", result)
}
