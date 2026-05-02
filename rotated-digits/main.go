package main

import (
	"fmt"
)

func rotatedDigits(n int) int {
	count := 0

	for i := 1; i <= n; i++ {
		if isGood(i) {
			count++
		}
	}

	return count
}

func isGood(num int) bool {
	changed := false

	for num > 0 {
		digit := num % 10

		switch digit {
		case 0, 1, 8:
			// Valid digits, remain same after rotation
		case 2, 5, 6, 9:
			// Valid digits, change after rotation
			changed = true
		default:
			// Invalid digits
			return false
		}

		num /= 10
	}

	return changed
}

func main() {
	var n int
	fmt.Print("Enter n: ")
	fmt.Scan(&n)

	result := rotatedDigits(n)

	fmt.Printf("Number of good integers between 1 and %d: %d\n", n, result)
}
