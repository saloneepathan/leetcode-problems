package main

import "fmt"

func minSumOfLengths(arr []int, target int) int {
	n := len(arr)

	const INF = int(1e9)

	// dp[i] = minimum length of a valid subarray
	// completely contained in arr[0...i].
	dp := make([]int, n)
	for i := range dp {
		dp[i] = INF
	}

	left := 0
	sum := 0
	ans := INF

	for right := 0; right < n; right++ {
		sum += arr[right]

		// Shrink window while sum is too large.
		for left <= right && sum > target {
			sum -= arr[left]
			left++
		}

		// If current window has sum == target.
		if sum == target {
			length := right - left + 1

			// Need a previous subarray that ends before
			// the current one starts.
			if left > 0 && dp[left-1] != INF {
				ans = min(ans, length+dp[left-1])
			}

			// This is the shortest valid subarray ending
			// at this position.
			if length < dp[right] {
				dp[right] = length
			}
		}

		// Carry forward the best previous subarray.
		if right > 0 {
			dp[right] = min(dp[right], dp[right-1])
		}
	}

	if ans == INF {
		return -1
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minSumOfLengths([]int{3, 2, 2, 4, 3}, 3))
	fmt.Println(minSumOfLengths([]int{7, 3, 4, 7}, 7))
	fmt.Println(minSumOfLengths([]int{4, 3, 2, 6, 2, 3, 4}, 6))
}
