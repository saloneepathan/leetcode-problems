package main

import "fmt"

func maxPalindromes(s string, k int) int {
	n := len(s)

	// isPalindrome[i][j] tells whether s[i:j+1] is a palindrome.
	isPalindrome := make([][]bool, n)
	for i := range isPalindrome {
		isPalindrome[i] = make([]bool, n)
	}

	// Build palindrome table.
	for length := 1; length <= n; length++ {
		for left := 0; left+length <= n; left++ {
			right := left + length - 1

			isPalindrome[left][right] = s[left] == s[right] &&
				(length <= 2 || isPalindrome[left+1][right-1])
		}
	}

	// dp[i] = maximum number of non-overlapping palindromes
	// using the first i characters.
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		// Don't use a palindrome ending at i-1.
		dp[i] = dp[i-1]

		// Try every palindrome of length >= k ending at i-1.
		for j := 0; j+k <= i; j++ {
			if isPalindrome[j][i-1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example 1
	s := "abaccdbbd"
	k := 3

	fmt.Println(maxPalindromes(s, k))

	// Add more test cases if needed.
	fmt.Println(maxPalindromes("adbcda", 2))
	fmt.Println(maxPalindromes("aaaaa", 2))
}
