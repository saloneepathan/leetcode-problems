package main

import "fmt"

func generateParenthesis(n int) []string {
	result := []string{}

	var backtrack func(string, int, int)

	backtrack = func(current string, open int, close int) {
		// A complete valid combination
		if len(current) == 2*n {
			result = append(result, current)
			return
		}

		// Add an opening parenthesis
		if open < n {
			backtrack(current+"(", open+1, close)
		}

		// Add a closing parenthesis only when valid
		if close < open {
			backtrack(current+")", open, close+1)
		}
	}

	backtrack("", 0, 0)

	return result
}

func main() {
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(1))
}