package main

import (
	"fmt"
	"strings"
)

func numOfStrings(patterns []string, word string) int {
	count := 0

	for _, pattern := range patterns {
		if strings.Contains(word, pattern) {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(numOfStrings([]string{"a", "abc", "bc", "d"}, "abc")) // 3
	fmt.Println(numOfStrings([]string{"a", "b", "c"}, "aaaaabbbbb"))  // 2
	fmt.Println(numOfStrings([]string{"a", "a", "a"}, "ab"))          // 3
}
