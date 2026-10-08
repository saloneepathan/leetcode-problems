package main

import "fmt"

func removeOuterParentheses(s string) string {
	var ans []rune
	depth := 0

	for _, c := range s {
		if c == '(' {
			// If depth > 0, this '(' is not an outermost '('
			if depth > 0 {
				ans = append(ans, c)
			}
			depth++
		} else if c == ')' {
			depth--

			// If depth > 0, this ')' is not an outermost ')'
			if depth > 0 {
				ans = append(ans, c)
			}
		}
	}

	return string(ans)
}

func main() {
	s := "(()())(())"

	result := removeOuterParentheses(s)

	fmt.Println(result)
}