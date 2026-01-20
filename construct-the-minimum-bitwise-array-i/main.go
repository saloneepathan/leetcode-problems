package main

import "fmt"

func minBitwiseArray(nums []int) []int {
	ans := make([]int, len(nums))

	for i, p := range nums {
		if p == 2 {
			ans[i] = -1
			continue
		}

		bit := 0
		for (p>>bit)&1 == 1 {
			bit++
		}

		ans[i] = p - (1 << (bit - 1))
	}

	return ans
}

func main() {
	nums := []int{2, 3, 5, 7}
	result := minBitwiseArray(nums)
	fmt.Println(result)
}
