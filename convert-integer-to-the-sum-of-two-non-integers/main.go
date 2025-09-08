package main

import (
	"fmt"
	"strconv"
)

// helper function to check if a number is a no-zero integer
func isNoZero(num int) bool {
	s := strconv.Itoa(num)
	for _, ch := range s {
		if ch == '0' {
			return false
		}
	}
	return true
}

func getNoZeroIntegers(n int) []int {
	for a := 1; a < n; a++ {
		b := n - a
		if isNoZero(a) && isNoZero(b) {
			return []int{a, b}
		}
	}
	return nil
}

func main() {
	fmt.Println(getNoZeroIntegers(2))   // [1 1]
	fmt.Println(getNoZeroIntegers(11))  // [2 9] or [8 3]
	fmt.Println(getNoZeroIntegers(101)) // [1 100] would fail, but e.g. [11 90] works
}
