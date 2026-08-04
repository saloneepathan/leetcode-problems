package main

import "fmt"

func maxProduct(n int) int {
	// Extract digits
	var digits []int
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	maxProd := 0

	// Check every pair
	for i := 0; i < len(digits); i++ {
		for j := i + 1; j < len(digits); j++ {
			product := digits[i] * digits[j]
			if product > maxProd {
				maxProd = product
			}
		}
	}

	return maxProd
}

func main() {
	fmt.Println(maxProduct(31))  // 3
	fmt.Println(maxProduct(22))  // 4
	fmt.Println(maxProduct(124)) // 8
	fmt.Println(maxProduct(987)) // 72
	fmt.Println(maxProduct(909)) // 81
}
