package main

import (
	"fmt"
)

func minimumPushes(word string) int {
	n := len(word)
	ans := 0

	// There are 8 available keys (2 to 9).
	// Every group of 8 characters requires one extra push.
	for i := 0; i < n; i++ {
		ans += (i / 8) + 1
	}

	return ans
}

func main() {
	fmt.Println(minimumPushes("abcde"))      // 5
	fmt.Println(minimumPushes("xycdefghij")) // 12
	fmt.Println(minimumPushes("a"))          // 1
	fmt.Println(minimumPushes("abcdefgh"))   // 8
	fmt.Println(minimumPushes("abcdefghi"))  // 10
}
