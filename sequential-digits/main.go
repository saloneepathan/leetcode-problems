package main

import "fmt"

func sequentialDigits(low int, high int) []int {
	result := []int{}
	digits := "123456789"

	// Generate all sequential digit numbers.
	for length := 2; length <= 9; length++ {
		for start := 0; start+length <= 9; start++ {
			num := 0
			for i := start; i < start+length; i++ {
				num = num*10 + int(digits[i]-'0')
			}

			if num >= low && num <= high {
				result = append(result, num)
			}
		}
	}

	return result
}

func main() {
	// Example 1
	low1, high1 := 100, 300
	fmt.Printf("Input: low = %d, high = %d\n", low1, high1)
	fmt.Println("Output:", sequentialDigits(low1, high1))

	// Example 2
	low2, high2 := 1000, 13000
	fmt.Printf("\nInput: low = %d, high = %d\n", low2, high2)
	fmt.Println("Output:", sequentialDigits(low2, high2))
}
