
package main

import (
	"fmt"
	"strings"
)

func maxSumOfSquares(num int, sum int) string {
	if sum > 9*num {
		return ""
	}

	var result strings.Builder
	result.Grow(num)

	for i := 0; i < num; i++ {
		digit := sum
		if digit > 9 {
			digit = 9
		}

		result.WriteByte(byte('0' + digit))
		sum -= digit
	}

	return result.String()
}

func main() {
	fmt.Println(maxSumOfSquares(2, 3))  // "30"
	fmt.Println(maxSumOfSquares(2, 17)) // "98"
	fmt.Println(maxSumOfSquares(1, 10)) // ""
	fmt.Println(maxSumOfSquares(3, 20)) // "992"
}
