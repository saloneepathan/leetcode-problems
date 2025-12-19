package main

import "fmt"

func maxProfit(prices []int, strategy []int, k int) int64 {
	n := len(prices)

	// Prefix sums
	profitSum := make([]int64, n+1)
	priceSum := make([]int64, n+1)

	for i := 0; i < n; i++ {
		profitSum[i+1] = profitSum[i] + int64(prices[i])*int64(strategy[i])
		priceSum[i+1] = priceSum[i] + int64(prices[i])
	}

	// Case: no modification
	res := profitSum[n]

	// Try every window of size k
	for i := k - 1; i < n; i++ {
		// profit before window
		leftProfit := profitSum[i-k+1]

		// profit after window
		rightProfit := profitSum[n] - profitSum[i+1]

		// profit gained by modifying:
		// first k/2 -> 0 (removes original profit)
		// last k/2 -> 1 (adds prices)
		changeProfit := priceSum[i+1] - priceSum[i-k/2+1]

		res = max(res, leftProfit+changeProfit+rightProfit)
	}

	return res
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example 1
	prices1 := []int{4, 2, 8}
	strategy1 := []int{-1, 0, 1}
	k1 := 2
	fmt.Println(maxProfit(prices1, strategy1, k1)) // Expected: 10

	// Example 2
	prices2 := []int{5, 4, 3}
	strategy2 := []int{1, 1, 0}
	k2 := 2
	fmt.Println(maxProfit(prices2, strategy2, k2)) // Expected: 9
}
