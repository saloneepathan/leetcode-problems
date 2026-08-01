package main

import "fmt"

func predictTheWinner(nums []int) bool {
	n := len(nums)

	memo := make([][]int, n)
	visited := make([][]bool, n)

	for i := 0; i < n; i++ {
		memo[i] = make([]int, n)
		visited[i] = make([]bool, n)
	}

	var dfs func(int, int) int

	dfs = func(left, right int) int {
		// Base case: only one number left
		if left == right {
			return nums[left]
		}

		if visited[left][right] {
			return memo[left][right]
		}

		// Choose left or right
		pickLeft := nums[left] - dfs(left+1, right)
		pickRight := nums[right] - dfs(left, right-1)

		if pickLeft > pickRight {
			memo[left][right] = pickLeft
		} else {
			memo[left][right] = pickRight
		}

		visited[left][right] = true
		return memo[left][right]
	}

	return dfs(0, n-1) >= 0
}

func main() {
	testCases := [][]int{
		{1, 5, 2},
		{1, 5, 233, 7},
		{1},
		{0},
		{1, 1},
		{2, 4, 55, 6, 8},
	}

	for _, nums := range testCases {
		fmt.Printf("nums = %v -> %v\n", nums, predictTheWinner(nums))
	}
}