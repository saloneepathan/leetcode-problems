package main

import (
	"fmt"
)

func countPermutations(complexity []int) int {
	n := len(complexity)

	// Check: all complexities except index 0 must be > complexity[0]
	for i := 1; i < n; i++ {
		if complexity[i] <= complexity[0] {
			return 0
		}
	}

	mod := 1000000007
	ans := 1

	// Multiply from 2 to n-1
	for i := 2; i < n; i++ {
		ans = ans * i % mod
	}

	return ans
}

func main() {
	fmt.Println(countPermutations([]int{1, 2, 3}))          // Example → 2
	fmt.Println(countPermutations([]int{3, 3, 3, 4, 4, 4})) // Example → 0

	// Additional tests
	fmt.Println(countPermutations([]int{5, 6, 7, 8})) // 6
	fmt.Println(countPermutations([]int{10, 9}))      // 0
}
