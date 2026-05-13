package main

import "fmt"

func minMoves(nums []int, limit int) int {
	n := len(nums)

	// Difference array
	diff := make([]int, 2*limit+2)

	for i := 0; i < n/2; i++ {
		a := nums[i]
		b := nums[n-1-i]

		low := min(a, b) + 1
		high := max(a, b) + limit
		sum := a + b

		/*
			For each pair:
			- Default = 2 moves
			- 1 move for sums in [low, high]
			- 0 move for exact sum
		*/

		// Add 2 moves for all possible sums
		diff[2] += 2

		// Reduce to 1 move in [low, high]
		diff[low] -= 1
		diff[high+1] += 1

		// Reduce to 0 move at exact sum
		diff[sum] -= 1
		diff[sum+1] += 1
	}

	ans := n
	cur := 0

	for s := 2; s <= 2*limit; s++ {
		cur += diff[s]
		ans = min(ans, cur)
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minMoves([]int{1, 2, 4, 3}, 4)) // Output: 1
	fmt.Println(minMoves([]int{1, 2, 2, 1}, 2)) // Output: 2
	fmt.Println(minMoves([]int{1, 2, 1, 2}, 2)) // Output: 0

	// Custom testcase
	fmt.Println(minMoves([]int{28, 50, 76, 80, 64, 30, 32, 84, 53, 8}, 84)) // Output: 4
}
