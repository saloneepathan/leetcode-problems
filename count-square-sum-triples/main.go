package main

import (
	"fmt"
	"math"
)

func countTriples(n int) int {
	count := 0

	for a := 1; a <= n; a++ {
		for b := 1; b <= n; b++ {
			c2 := a*a + b*b
			c := int(math.Sqrt(float64(c2)))
			if c*c == c2 && c <= n {
				count++
			}
		}
	}

	return count
}

func main() {
	fmt.Println(countTriples(5))  // Output: 2
	fmt.Println(countTriples(10)) // Output: 4
}
