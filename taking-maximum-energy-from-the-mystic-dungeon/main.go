package main

import "fmt"

func maximumEnergy(energy []int, k int) int {
	n := len(energy)
	dp := make([]int, n)
	maxEnergy := -1000000000 // very small initial value

	for i := n - 1; i >= 0; i-- {
		dp[i] = energy[i]
		if i+k < n {
			dp[i] += dp[i+k]
		}
		if dp[i] > maxEnergy {
			maxEnergy = dp[i]
		}
	}
	return maxEnergy
}

func main() {
	fmt.Println(maximumEnergy([]int{5, 2, -10, -5, 1}, 3)) // Output: 3
	fmt.Println(maximumEnergy([]int{-2, -3, -1}, 2))       // Output: -1
}
