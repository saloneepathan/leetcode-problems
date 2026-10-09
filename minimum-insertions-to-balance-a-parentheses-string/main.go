package main

import "fmt"

func minInsertions(s string) int {
	insertions := 0
	open := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			open++
		} else {
			// If the next character is also ')',
			// we have a pair of closing parentheses.
			if i+1 < len(s) && s[i+1] == ')' {
				i++
			} else {
				// Insert one ')' to complete the pair.
				insertions++
			}

			if open > 0 {
				open--
			} else {
				// No opening '(' available; insert one.
				insertions++
			}
		}
	}

	// Every remaining '(' needs two closing parentheses.
	insertions += open * 2

	return insertions
}

func main() {
	fmt.Println(minInsertions("(()))"))  // 1
	fmt.Println(minInsertions("())"))    // 0
	fmt.Println(minInsertions("))())(")) // 3
}
