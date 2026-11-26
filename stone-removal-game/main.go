package main

import "fmt"

func canAliceWin(n int) bool {
	take, k := 10, 0

	for n >= take {
		n -= take
		take--
		k++
	}

	return k%2 == 1
}

func main() {
	testCases := []int{
		10,
		11,
		20,
		21,
		30,
		55,
		100,
	}

	for _, n := range testCases {
		fmt.Printf("n = %d -> Alice wins: %v\n", n, canAliceWin(n))
	}
}