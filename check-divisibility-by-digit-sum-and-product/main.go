package main

import "fmt"

func checkDivisibility(n int) bool {
	original := n
	digitSum := 0
	digitProduct := 1

	for n > 0 {
		digit := n % 10
		digitSum += digit
		digitProduct *= digit
		n /= 10
	}

	return original%(digitSum+digitProduct) == 0
}

func main() {
	fmt.Println(checkDivisibility(99)) // true
	fmt.Println(checkDivisibility(23)) // false
}