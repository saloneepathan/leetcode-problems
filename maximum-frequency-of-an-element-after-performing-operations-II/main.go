package main

import (
	"fmt"
	"sort"
)

func maxFrequency(nums []int, k int, numOperations int) int {
	sort.Ints(nums)
	n := len(nums)

	left := 0
	total := 0
	res := 1

	for right := 0; right < n; right++ {
		// Add current value into window sum
		total += nums[right]

		// while cost to make all nums in window equal to nums[right] > (k * numOperations)
		// Actually, here each element can move in [-k, k], so effective reachable range expands by 2*k
		for nums[right]-nums[left] > 2*k {
			total -= nums[left]
			left++
		}

		windowSize := right - left + 1

		// We can only perform at most numOperations changes
		if windowSize > numOperations+1 {
			windowSize = numOperations + 1
		}

		if windowSize > res {
			res = windowSize
		}
	}

	return res
}

func main() {
	fmt.Println(maxFrequency([]int{1, 4, 5}, 1, 2))       // 2
	fmt.Println(maxFrequency([]int{5, 11, 20, 20}, 5, 1)) // 2
	fmt.Println(maxFrequency([]int{94, 10}, 51, 1))       // 1
	fmt.Println(maxFrequency([]int{2, 2, 2}, 3, 2))       // 3
	fmt.Println(maxFrequency([]int{10, 100, 200}, 90, 2)) // 2
}
