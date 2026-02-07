package main

import (
	"fmt"
)

func minimumDeletions(s string) int {
	countB := 0
	deletions := 0

	for _, ch := range s {
		if ch == 'b' {
			countB++
		} else { // ch == 'a'
			// Either delete this 'a' or delete all previous 'b's
			if deletions+1 < countB {
				deletions = deletions + 1
			} else {
				deletions = countB
			}
		}
	}
	return deletions
}

func main() {
	// Test cases
	tests := []string{
		"aababbab",
		"bbaaaaabb",
		"a",
		"b",
		"abba",
		"bbbbbaaaaa",
	}

	for _, s := range tests {
		fmt.Printf("Input: %s -> Minimum Deletions: %d\n", s, minimumDeletions(s))
	}
}
