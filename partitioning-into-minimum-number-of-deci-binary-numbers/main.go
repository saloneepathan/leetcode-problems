package main

import (
	"fmt"
)

func minPartitions(n string) int {
	maxDigit := 0

	for _, ch := range n {
		digit := int(ch - '0')
		if digit > maxDigit {
			maxDigit = digit
		}

		// Early exit since max possible digit is 9
		if maxDigit == 9 {
			return 9
		}
	}

	return maxDigit
}

func main() {
	// Sample test cases
	tests := []string{
		"32",
		"82734",
		"27346209830709182346",
	}

	for _, test := range tests {
		fmt.Printf("Input: %s\nOutput: %d\n\n", test, minPartitions(test))
	}
}
