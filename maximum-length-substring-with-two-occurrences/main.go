package main

import "fmt"

func maximumLengthSubstring(s string) int {
	count := make([]int, 26)
	left := 0
	maxLen := 0

	for right := 0; right < len(s); right++ {
		count[s[right]-'a']++

		// Ensure every character occurs at most twice.
		for count[s[right]-'a'] > 2 {
			count[s[left]-'a']--
			left++
		}

		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}

func main() {
	s1 := "bcbbbcba"
	s2 := "aaaa"

	fmt.Println(maximumLengthSubstring(s1)) // 4
	fmt.Println(maximumLengthSubstring(s2)) // 2
}