package main

import (
	"fmt"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func minOperations(nums []int) int {
	n := len(nums)
	overallGcd := nums[0]
	for i := 1; i < n; i++ {
		overallGcd = gcd(overallGcd, nums[i])
	}
	if overallGcd != 1 {
		return -1
	}

	// If any element is already 1, just count how many are not 1
	ones := 0
	for _, v := range nums {
		if v == 1 {
			ones++
		}
	}
	if ones > 0 {
		return n - ones
	}

	// Otherwise, find the shortest subarray with gcd == 1
	minLen := n + 1
	for i := 0; i < n; i++ {
		currGcd := nums[i]
		for j := i + 1; j < n; j++ {
			currGcd = gcd(currGcd, nums[j])
			if currGcd == 1 {
				if j-i+1 < minLen {
					minLen = j - i + 1
				}
				break
			}
		}
	}

	// total operations = (length to make one '1') + (spread that 1 to others)
	return n + minLen - 2
}

func main() {
	fmt.Println(minOperations([]int{2, 6, 3, 4}))   // Output: 4
	fmt.Println(minOperations([]int{2, 10, 6, 14})) // Output: -1
}
