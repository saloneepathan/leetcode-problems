package main

import (
	"fmt"
)

func processStr(s string, k int64) byte {
	n := len(s)

	// lengths[i] = length of result after processing first i characters
	lengths := make([]int64, n+1)

	for i := 0; i < n; i++ {
		cur := lengths[i]

		switch s[i] {
		case '*':
			if cur > 0 {
				lengths[i+1] = cur - 1
			} else {
				lengths[i+1] = 0
			}
		case '#':
			lengths[i+1] = cur * 2
		case '%':
			lengths[i+1] = cur
		default: // lowercase letter
			lengths[i+1] = cur + 1
		}
	}

	// k out of bounds
	if k >= lengths[n] {
		return '.'
	}

	// Walk backwards to trace where index k came from
	for i := n - 1; i >= 0; i-- {
		ch := s[i]

		switch ch {
		case '%':
			// reverse
			k = lengths[i] - 1 - k

		case '#':
			// duplicated string
			if lengths[i] > 0 && k >= lengths[i] {
				k -= lengths[i]
			}

		case '*':
			// deletion of last character; valid indices remain unchanged

		default: // letter appended
			if k == lengths[i] {
				return byte(ch)
			}
		}
	}

	return '.'
}

func main() {
	fmt.Printf("%c\n", processStr("a#b%*", 1))  // a
	fmt.Printf("%c\n", processStr("cd%#*#", 3)) // d
	fmt.Printf("%c\n", processStr("z*#", 0))    // .
}
