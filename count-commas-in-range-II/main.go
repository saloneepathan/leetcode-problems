package main

import "fmt"

func countCommas(n int64) int64 {
	var count int64

	for place := int64(1000); place <= n; {
		count += n - place + 1

		// Prevent int64 overflow when multiplying by 1000.
		if place > n/1000 {
			break
		}

		place *= 1000
	}

	return count
}

func main() {
	var n int64

	fmt.Print("Enter n: ")
	fmt.Scan(&n)

	fmt.Println("Total commas:", countCommas(n))
}