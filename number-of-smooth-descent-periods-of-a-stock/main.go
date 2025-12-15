package main

import (
	"fmt"
)

func getDescentPeriods(prices []int) int64 {
	var count int64 = 1
	var streak int64 = 1

	for i := 1; i < len(prices); i++ {
		if prices[i-1]-prices[i] == 1 {
			streak++
		} else {
			streak = 1
		}
		count += streak
	}

	return count
}

func main() {
	// Example test cases
	prices1 := []int{3, 2, 1, 4}
	prices2 := []int{8, 6, 7, 7}
	prices3 := []int{1}

	fmt.Println(getDescentPeriods(prices1)) // Output: 7
	fmt.Println(getDescentPeriods(prices2)) // Output: 4
	fmt.Println(getDescentPeriods(prices3)) // Output: 1
}
