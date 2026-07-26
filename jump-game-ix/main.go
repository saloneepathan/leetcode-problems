package main

import "fmt"

func maxValue(nums []int) []int {
	n := len(nums)

	if n == 1 {
		return []int{nums[0]}
	}

	// suffixMin[i] = minimum value in nums[i...n-1]
	suffixMin := make([]int, n)
	suffixMin[n-1] = nums[n-1]

	for i := n - 2; i >= 0; i-- {
		if nums[i] < suffixMin[i+1] {
			suffixMin[i] = nums[i]
		} else {
			suffixMin[i] = suffixMin[i+1]
		}
	}

	ans := make([]int, n)

	start := 0
	currMax := nums[0]

	for i := 0; i < n-1; i++ {
		if nums[i] > currMax {
			currMax = nums[i]
		}

		// If max(left) <= min(right),
		// then there is NO inversion crossing this boundary,
		// so components split here.
		if currMax <= suffixMin[i+1] {
			for j := start; j <= i; j++ {
				ans[j] = currMax
			}

			start = i + 1
			currMax = nums[start]
		}
	}

	// Fill last component
	for i := start; i < n; i++ {
		ans[i] = currMax
	}

	return ans
}

func main() {
	fmt.Println(maxValue([]int{2, 1, 3}))    // [2 2 3]
	fmt.Println(maxValue([]int{2, 3, 1}))    // [3 3 3]
	fmt.Println(maxValue([]int{1, 3, 2}))    // [1 3 3]
	fmt.Println(maxValue([]int{2, 4, 1, 3})) // [4 4 4 4]
}
