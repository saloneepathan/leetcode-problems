package main

import (
	"fmt"
)

func numberOfSpecialChars(word string) int {
	lower := make(map[rune]bool)
	upper := make(map[rune]bool)

	// Store lowercase and uppercase letters
	for _, ch := range word {
		if ch >= 'a' && ch <= 'z' {
			lower[ch] = true
		} else if ch >= 'A' && ch <= 'Z' {
			upper[ch] = true
		}
	}

	count := 0

	// Check if both lowercase and uppercase exist
	for ch := 'a'; ch <= 'z'; ch++ {
		if lower[ch] && upper[ch-'a'+'A'] {
			count++
		}
	}

	return count
}

func main() {
	// Test cases
	fmt.Println(numberOfSpecialChars("aaAbcBC")) // Output: 3
	fmt.Println(numberOfSpecialChars("abc"))     // Output: 0
	fmt.Println(numberOfSpecialChars("abBCab"))  // Output: 1
}
