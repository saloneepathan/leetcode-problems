package main

import "fmt"

func separateDigits(nums []int) []int {
	var result []int

	for _, num := range nums {
		digits := []int{}

		// Extract digits in reverse order
		for num > 0 {
			digits = append(digits, num%10)
			num /= 10
		}

		// Add digits in correct order
		for i := len(digits) - 1; i >= 0; i-- {
			result = append(result, digits[i])
		}
	}

	return result
}

func main() {
	nums1 := []int{13, 25, 83, 77}
	fmt.Println(separateDigits(nums1)) // [1 3 2 5 8 3 7 7]

	nums2 := []int{7, 1, 3, 9}
	fmt.Println(separateDigits(nums2)) // [7 1 3 9]
}
