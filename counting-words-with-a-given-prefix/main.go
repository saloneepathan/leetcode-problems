package main

import (
	"fmt"
	"strings"
)

// prefixCount returns how many words start with pref.
func prefixCount(words []string, pref string) int {
	count := 0
	for _, w := range words {
		if strings.HasPrefix(w, pref) {
			count++
		}
	}
	return count
}

func main() {
	tests := []struct {
		words    []string
		pref     string
		expected int
	}{
		{[]string{"pay", "attention", "practice", "attend"}, "at", 2},
		{[]string{"leetcode", "win", "loops", "success"}, "code", 0},
		{[]string{"a"}, "a", 1},                        // single exact match
		{[]string{"a", "ab"}, "abc", 0},                // prefix longer than words
		{[]string{"abc", "abc", "abd"}, "abc", 2},      // duplicates count separately
		{[]string{"apple", "app", "ap", "a"}, "ap", 3}, // word equal to pref counts
	}
	for _, t := range tests {
		got := prefixCount(t.words, t.pref)
		status := "OK"
		if got != t.expected {
			status = "FAIL"
		}
		fmt.Printf("words=%q pref=%q expected=%d actual=%d %s\n", t.words, t.pref, t.expected, got, status)
	}
}
