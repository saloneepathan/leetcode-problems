package main

import (
	"fmt"
)

// maxOperations returns the maximum number of operations that can be performed
func maxOperations(s string) int {
	var ones int64 = 0
	var ans int64 = 0
	n := len(s)

	for i := 0; i < n; i++ {
		if s[i] == '1' {
			ones++
		} else {
			// If this '0' is the last zero in a block of zeros (i.e., next is '1' or end of string)
			if i+1 == n || s[i+1] == '1' {
				ans += ones
			}
		}
	}
	return int(ans)
}

func main() {
	fmt.Println(maxOperations("1001101")) // Expected: 4
	fmt.Println(maxOperations("00111"))   // Expected: 0
	fmt.Println(maxOperations("101010"))  // Additional test: Expected 6
}
