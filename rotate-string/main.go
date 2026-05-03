package main

import (
	"fmt"
	"strings"
)

// rotateString returns true if s can be rotated to become goal
func rotateString(s string, goal string) bool {
	// Lengths must match
	if len(s) != len(goal) {
		return false
	}

	// If goal is a substring of s+s, rotation is possible
	return strings.Contains(s+s, goal)
}

func main() {
	// Test cases
	fmt.Println(rotateString("abcde", "cdeab")) // true
	fmt.Println(rotateString("abcde", "abced")) // false
	fmt.Println(rotateString("aa", "aa"))       // true
}
