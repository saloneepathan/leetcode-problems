package main

import (
	"fmt"
	"strings"
)

func canBeTypedWords(text string, brokenLetters string) int {
	// Step 1: Build a lookup table for broken letters
	var broken [26]bool
	for _, ch := range brokenLetters {
		broken[ch-'a'] = true
	}

	// Step 2: Split text into words
	words := strings.Split(text, " ")
	count := 0

	// Step 3: Check each word
	for _, word := range words {
		canType := true
		for _, ch := range word {
			if broken[ch-'a'] { // O(1) lookup
				canType = false
				break // stop early for performance
			}
		}
		if canType {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(canBeTypedWords("hello world", "ad")) // 1
	fmt.Println(canBeTypedWords("leet code", "lt"))   // 1
	fmt.Println(canBeTypedWords("leet code", "e"))    // 0
}
