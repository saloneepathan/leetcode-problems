package main

import (
	"fmt"
)

func reverseBits(n int) int {
	result := 0

	for i := 0; i < 32; i++ {
		result <<= 1
		result |= (n & 1)
		n >>= 1
	}

	return result
}

func main() {

	// Example 1
	n1 := 43261596
	fmt.Println("Input:", n1)
	fmt.Println("Output:", reverseBits(n1))

	// Example 2
	n2 := 2147483644
	fmt.Println("Input:", n2)
	fmt.Println("Output:", reverseBits(n2))

}
