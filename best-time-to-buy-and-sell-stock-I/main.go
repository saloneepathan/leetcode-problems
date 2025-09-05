package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	minPrice := math.MaxInt32
	maxProfit := 0

	for _, price := range prices {
		// update min price if current price is lower
		if price < minPrice {
			minPrice = price
		}
		// calculate profit if sold today
		profit := price - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}

func main() {
	// Example 1
	prices1 := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("Max Profit:", maxProfit(prices1)) // Output: 5

	// Example 2
	prices2 := []int{7, 6, 4, 3, 1}
	fmt.Println("Max Profit:", maxProfit(prices2)) // Output: 0
}
