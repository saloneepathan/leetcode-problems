package main

import "fmt"

func smallestNumber(n int, t int) int {
	for {
		if digitProduct(n)%t == 0 {
			return n
		}
		n++
	}
}

func digitProduct(num int) int {
	product := 1

	for num > 0 {
		digit := num % 10
		if digit == 0 {
			return 0
		}
		product *= digit
		num /= 10
	}

	return product
}

func main() {
	// Test Case 1
	n1, t1 := 10, 2
	fmt.Printf("Input: n = %d, t = %d\n", n1, t1)
	fmt.Printf("Output: %d\n\n", smallestNumber(n1, t1))

	// Test Case 2
	n2, t2 := 15, 3
	fmt.Printf("Input: n = %d, t = %d\n", n2, t2)
	fmt.Printf("Output: %d\n\n", smallestNumber(n2, t2))

	// Additional Test Cases
	tests := [][]int{
		{1, 1},
		{19, 2},
		{25, 5},
		{99, 10},
		{100, 7},
	}

	for _, test := range tests {
		n, t := test[0], test[1]
		fmt.Printf("Input: n = %d, t = %d -> Output: %d\n",
			n, t, smallestNumber(n, t))
	}
}
