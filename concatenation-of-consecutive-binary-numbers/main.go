package main

import (
	"fmt"
)

func concatenatedBinary(n int) int {
	const mod = 1000000007
	result := 0
	bitLength := 0

	for i := 1; i <= n; i++ {
		// If i is power of 2, increase bit length
		if (i & (i - 1)) == 0 {
			bitLength++
		}

		// Left shift result and add i
		result = ((result << bitLength) | i) % mod
	}

	return result
}

func main() {
	fmt.Println(concatenatedBinary(1))  // Output: 1
	fmt.Println(concatenatedBinary(3))  // Output: 27
	fmt.Println(concatenatedBinary(12)) // Output: 505379714
}
