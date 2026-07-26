package main

import "fmt"

func countOperations(num1 int, num2 int) int {
	ops := 0
	for num1 != 0 && num2 != 0 {
		if num1 >= num2 {
			ops += num1 / num2
			num1 = num1 % num2
		} else {
			ops += num2 / num1
			num2 = num2 % num1
		}
	}
	return ops
}

func main() {
	// Example 1
	num1, num2 := 2, 3
	fmt.Printf("Input: num1 = %d, num2 = %d → Output: %d\n", num1, num2, countOperations(num1, num2))

	// Example 2
	num1, num2 = 10, 10
	fmt.Printf("Input: num1 = %d, num2 = %d → Output: %d\n", num1, num2, countOperations(num1, num2))

	// Additional tests
	num1, num2 = 5, 4
	fmt.Printf("Input: num1 = %d, num2 = %d → Output: %d\n", num1, num2, countOperations(num1, num2))

	num1, num2 = 100, 25
	fmt.Printf("Input: num1 = %d, num2 = %d → Output: %d\n", num1, num2, countOperations(num1, num2))
}
