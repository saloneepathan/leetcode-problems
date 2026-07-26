package main

import (
	"fmt"
)

func maxJumps(arr []int, d int) int {
	n := len(arr)
	f := make([]int, n)

	// Initialize memoization array
	for i := range f {
		f[i] = -1
	}

	var dfs func(int)

	dfs = func(id int) {
		// Already computed
		if f[id] != -1 {
			return
		}

		// Minimum jumps count is 1 (the current index itself)
		f[id] = 1

		// Check left side
		for i := id - 1; i >= 0 && id-i <= d; i-- {
			// Stop if a higher or equal element blocks the jump
			if arr[i] >= arr[id] {
				break
			}

			dfs(i)

			if f[i]+1 > f[id] {
				f[id] = f[i] + 1
			}
		}

		// Check right side
		for i := id + 1; i < n && i-id <= d; i++ {
			// Stop if a higher or equal element blocks the jump
			if arr[i] >= arr[id] {
				break
			}

			dfs(i)

			if f[i]+1 > f[id] {
				f[id] = f[i] + 1
			}
		}
	}

	// Compute result for every index
	for i := 0; i < n; i++ {
		dfs(i)
	}

	// Find maximum value
	ans := 0
	for _, val := range f {
		if val > ans {
			ans = val
		}
	}

	return ans
}

func main() {
	arr := []int{6, 4, 14, 6, 8, 13, 9, 7, 10, 6, 12}
	d := 2

	result := maxJumps(arr, d)

	fmt.Println("Maximum jumps:", result)
}
