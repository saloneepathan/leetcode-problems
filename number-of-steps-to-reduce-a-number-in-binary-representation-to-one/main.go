package main

import (
	"fmt"
)

func numSteps(s string) int {
	steps := 0
	carry := 0
	n := len(s)

	// Traverse from right to left (excluding the first bit)
	for i := n - 1; i > 0; i-- {
		bit := int(s[i] - '0')

		if bit+carry == 1 {
			// Odd → add 1 (creates carry), then divide
			steps += 2
			carry = 1
		} else {
			// Even → just divide by 2
			steps += 1
		}
	}

	// If there's leftover carry, add one more step
	return steps + carry
}

func main() {
	// Testcases
	fmt.Println(numSteps("1101")) // Output: 6
	fmt.Println(numSteps("10"))   // Output: 1
	fmt.Println(numSteps("1"))    // Output: 0
}
