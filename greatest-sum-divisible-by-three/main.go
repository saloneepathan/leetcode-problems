package main

import (
	"fmt"
)

func maxSumDivThree(nums []int) int {
	total := 0

	// smallest two numbers with remainder 1 and 2
	rem1 := []int{10001, 10001}
	rem2 := []int{10001, 10001}

	for _, num := range nums {
		total += num

		if num%3 == 1 {
			if num < rem1[0] {
				rem1[1] = rem1[0]
				rem1[0] = num
			} else if num < rem1[1] {
				rem1[1] = num
			}
		} else if num%3 == 2 {
			if num < rem2[0] {
				rem2[1] = rem2[0]
				rem2[0] = num
			} else if num < rem2[1] {
				rem2[1] = num
			}
		}
	}

	if total%3 == 0 {
		return total
	}

	if total%3 == 1 {
		option1 := rem1[0]
		option2 := rem2[0] + rem2[1]
		remove := minValid(option1, option2)
		return total - remove
	}

	// remainder = 2
	option1 := rem2[0]
	option2 := rem1[0] + rem1[1]
	remove := minValid(option1, option2)
	return total - remove
}

func minValid(a, b int) int {
	if a >= 10001 && b >= 10001 {
		return 0
	}
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxSumDivThree([]int{3, 6, 5, 1, 8})) // 18
	fmt.Println(maxSumDivThree([]int{4}))             // 0
	fmt.Println(maxSumDivThree([]int{1, 2, 3, 4, 4})) // 12
}
