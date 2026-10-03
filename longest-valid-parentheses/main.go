package main

import "fmt"

func longestValidParentheses(s string) int {
	stack := []int{-1}
	maxLen := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]

			if len(stack) == 0 {
				// Current ')' cannot be part of a valid substring.
				stack = append(stack, i)
			} else {
				// Everything after stack top forms a valid substring.
				length := i - stack[len(stack)-1]
				if length > maxLen {
					maxLen = length
				}
			}
		}
	}

	return maxLen
}

func main() {
	fmt.Println(longestValidParentheses("(()"))
	fmt.Println(longestValidParentheses(")()())"))
	fmt.Println(longestValidParentheses(""))
}
