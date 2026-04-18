package main

import (
	"fmt"
)

func mirrorDistance(n int) int {
	original := n
	reversed := 0

	for n > 0 {
		digit := n % 10
		reversed = reversed*10 + digit
		n /= 10
	}

	if original > reversed {
		return original - reversed
	}
	return reversed - original
}

func main() {
	// Test cases
	fmt.Println(mirrorDistance(25)) // Output: 27
	fmt.Println(mirrorDistance(10)) // Output: 9
	fmt.Println(mirrorDistance(7))  // Output: 0

	// You can also take input from user if needed
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)
	fmt.Println("Mirror Distance:", mirrorDistance(n))
}
