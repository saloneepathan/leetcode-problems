package main

import (
	"fmt"
)

func reverseParentheses(s string) string {
	stack := []string{""}

	for _, ch := range s {
		if ch == '(' {
			// Start a new substring
			stack = append(stack, "")
		} else if ch == ')' {
			// Get the current substring
			curr := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// Reverse it
			runes := []rune(curr)
			for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
				runes[i], runes[j] = runes[j], runes[i]
			}

			// Append reversed substring to previous level
			stack[len(stack)-1] += string(runes)
		} else {
			// Add character to current substring
			stack[len(stack)-1] += string(ch)
		}
	}

	return stack[0]
}

func main() {
	fmt.Println(reverseParentheses("(abcd)"))         // dcba
	fmt.Println(reverseParentheses("(u(love)i)"))     // iloveu
	fmt.Println(reverseParentheses("(ed(et(oc))el)")) // leetcode
}
