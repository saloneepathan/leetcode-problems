package main

import "fmt"

func sumAndMultiply(n int) int64 {
	if n == 0 {
		return 0
	}

	div := 1
	for div <= n/10 {
		div *= 10
	}

	var x int64
	sum := 0

	for div > 0 {
		digit := n / div
		if digit != 0 {
			x = x*10 + int64(digit)
			sum += digit
		}
		n %= div
		div /= 10
	}

	return x * int64(sum)
}

func main() {
	testCases := []int{
		10203004,
		1000,
		0,
		12345,
		9000009,
	}

	for _, n := range testCases {
		fmt.Printf("Input: %d -> Output: %d\n", n, sumAndMultiply(n))
	}
}
