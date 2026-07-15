package main

import (
	"fmt"
)

func gcdOfOddEvenSums(n int) int {
	// Mathematical proof:
	// Sum of first n odd numbers = n^2
	// Sum of first n even numbers = n*(n+1)
	// GCD(n^2, n*(n+1)) = n
	return n
}

func main() {
	// Test cases
	fmt.Println(gcdOfOddEvenSums(4))  // Output: 4
	fmt.Println(gcdOfOddEvenSums(5))  // Output: 5
	fmt.Println(gcdOfOddEvenSums(1))  // Output: 1
	fmt.Println(gcdOfOddEvenSums(10)) // Output: 10
}
