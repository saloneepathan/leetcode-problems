package main

import (
	"fmt"
)

func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}

	mask := 0
	temp := n

	for temp > 0 {
		mask = (mask << 1) | 1
		temp >>= 1
	}

	return n ^ mask
}

func main() {
	fmt.Println(bitwiseComplement(5))  // Output: 2
	fmt.Println(bitwiseComplement(7))  // Output: 0
	fmt.Println(bitwiseComplement(10)) // Output: 5
}
