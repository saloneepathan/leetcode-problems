package main

import (
	"fmt"
)

func numberOfSpecialChars(word string) int {
	// last occurrence index of lowercase letters
	lastLower := make([]int, 26)

	// first occurrence index of uppercase letters
	firstUpper := make([]int, 26)

	for i := 0; i < 26; i++ {
		lastLower[i] = -1
		firstUpper[i] = -1
	}

	// Traverse the string
	for i, ch := range word {
		if ch >= 'a' && ch <= 'z' {
			lastLower[ch-'a'] = i
		} else {
			idx := ch - 'A'
			if firstUpper[idx] == -1 {
				firstUpper[idx] = i
			}
		}
	}

	// Count special characters
	count := 0

	for i := 0; i < 26; i++ {
		if lastLower[i] != -1 &&
			firstUpper[i] != -1 &&
			lastLower[i] < firstUpper[i] {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(numberOfSpecialChars("aaAbcBC")) // 3
	fmt.Println(numberOfSpecialChars("abc"))     // 0
	fmt.Println(numberOfSpecialChars("AbBCab"))  // 0
}
