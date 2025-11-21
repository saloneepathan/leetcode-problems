package main

import (
	"fmt"
)

func countPalindromicSubsequence(s string) int {
	first := make([]int, 26)
	last := make([]int, 26)

	// Initialize with invalid indices
	for i := 0; i < 26; i++ {
		first[i] = -1
		last[i] = -1
	}

	// Record first and last occurrences of each character
	for i := 0; i < len(s); i++ {
		idx := int(s[i] - 'a')
		if first[idx] == -1 {
			first[idx] = i
		}
		last[idx] = i
	}

	result := 0

	// For each character, check unique chars between first and last index
	for c := 0; c < 26; c++ {
		if first[c] != -1 && last[c] != -1 && first[c] < last[c] {
			seen := make(map[byte]bool)

			for i := first[c] + 1; i < last[c]; i++ {
				seen[s[i]] = true
			}

			result += len(seen)
		}
	}

	return result
}

func main() {
	fmt.Println(countPalindromicSubsequence("aabca"))   // Output: 3
	fmt.Println(countPalindromicSubsequence("adc"))     // Output: 0
	fmt.Println(countPalindromicSubsequence("bbcbaba")) // Output: 4
}
