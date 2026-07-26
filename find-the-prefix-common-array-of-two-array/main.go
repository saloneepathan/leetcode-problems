package main

import (
	"fmt"
)

func findThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)

	freq := make([]int, n+1)
	result := make([]int, n)

	common := 0

	for i := 0; i < n; i++ {
		freq[A[i]]++
		if freq[A[i]] == 2 {
			common++
		}

		freq[B[i]]++
		if freq[B[i]] == 2 {
			common++
		}

		result[i] = common
	}

	return result
}

func main() {
	A1 := []int{1, 3, 2, 4}
	B1 := []int{3, 1, 2, 4}

	fmt.Println(findThePrefixCommonArray(A1, B1))
	// Output: [0 2 3 4]

	A2 := []int{2, 3, 1}
	B2 := []int{3, 1, 2}

	fmt.Println(findThePrefixCommonArray(A2, B2))
	// Output: [0 1 3]
}
