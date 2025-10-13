package main

import (
	"fmt"
	"strings"
)

// helper: generate frequency-based signature for a word
func signature(s string) string {
	count := [26]int{}
	for _, ch := range s {
		count[ch-'a']++
	}

	// convert counts to a compact string key
	var sb strings.Builder
	for i := 0; i < 26; i++ {
		if count[i] > 0 {
			sb.WriteByte(byte('a' + i))
			sb.WriteString(fmt.Sprint(count[i]))
		}
	}
	return sb.String()
}

func removeAnagrams(words []string) []string {
	stack := []string{}
	lastSig := ""

	for _, word := range words {
		sig := signature(word)
		if sig != lastSig {
			stack = append(stack, word)
			lastSig = sig
		}
	}

	return stack
}

func main() {
	fmt.Println(removeAnagrams([]string{"abba", "baba", "bbaa", "cd", "cd"})) // ["abba", "cd"]
	fmt.Println(removeAnagrams([]string{"a", "b", "c", "d", "e"}))            // ["a", "b", "c", "d", "e"]
}
