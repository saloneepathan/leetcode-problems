package main

import "fmt"

func checkValidString(s string) bool {
	low, high := 0, 0

	for _, ch := range s {
		switch ch {
		case '(':
			low++
			high++

		case ')':
			low--
			high--

		case '*':
			// '*' can be ')', '(' or empty.
			low--
			high++
		}

		// We cannot have fewer than 0 unmatched '('.
		if low < 0 {
			low = 0
		}

		// Even the maximum possible '(' count is negative.
		if high < 0 {
			return false
		}
	}

	return low == 0
}

func main() {
	fmt.Println(checkValidString("()"))   // true
	fmt.Println(checkValidString("(*)"))  // true
	fmt.Println(checkValidString("(*))")) // true
	fmt.Println(checkValidString("("))    // false
}