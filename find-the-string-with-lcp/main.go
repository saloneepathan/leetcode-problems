package main

import (
	"fmt"
)

func findTheString(lcp [][]int) string {
	n := len(lcp)
	word := make([]byte, n)
	current := byte('a')

	// Step 1: Construct the lexicographically smallest string
	for i := 0; i < n; i++ {
		if word[i] == 0 {
			if current > 'z' {
				return ""
			}
			word[i] = current

			for j := i + 1; j < n; j++ {
				if lcp[i][j] > 0 {
					word[j] = word[i]
				}
			}
			current++
		}
	}

	// Step 2: Verify LCP matrix
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if word[i] == word[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = 0
			}

			if dp[i][j] != lcp[i][j] {
				return ""
			}
		}
	}

	return string(word)
}

func main() {

	test1 := [][]int{
		{4, 0, 2, 0},
		{0, 3, 0, 1},
		{2, 0, 2, 0},
		{0, 1, 0, 1},
	}

	test2 := [][]int{
		{4, 3, 2, 1},
		{3, 3, 2, 1},
		{2, 2, 2, 1},
		{1, 1, 1, 1},
	}

	test3 := [][]int{
		{4, 3, 2, 1},
		{3, 3, 2, 1},
		{2, 2, 2, 1},
		{1, 1, 1, 3},
	}

	fmt.Println(findTheString(test1)) // abab
	fmt.Println(findTheString(test2)) // aaaa
	fmt.Println(findTheString(test3)) // ""
}
