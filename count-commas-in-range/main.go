package main

import "fmt"

func countCommas(n int) int {
	count := 0

	for i := 1000; i <= n; i++ {
		x := i

		for x >= 1000 {
			count++
			x /= 1000
		}
	}

	return count
}

func main() {
	n := 1002

	fmt.Println(countCommas(n))
}
