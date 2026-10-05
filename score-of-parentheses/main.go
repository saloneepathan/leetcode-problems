package main

import "fmt"

func scoreOfParentheses(s string) int {
	stack := []int{0}

	for _, ch := range s {
		if ch == '(' {
			// Start a new nested score
			stack = append(stack, 0)
		} else {
			// Get score inside current pair
			inner := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			score := 0
			if inner == 0 {
				// ()
				score = 1
			} else {
				// (A)
				score = 2 * inner
			}

			// Add this score to the parent level
			stack[len(stack)-1] += score
		}
	}

	return stack[0]
}

func main() {
	fmt.Println(scoreOfParentheses("()"))     // 1
	fmt.Println(scoreOfParentheses("(())"))   // 2
	fmt.Println(scoreOfParentheses("()()"))   // 2
	fmt.Println(scoreOfParentheses("(()())")) // 4
}
