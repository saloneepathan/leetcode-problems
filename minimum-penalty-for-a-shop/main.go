package main

import (
	"fmt"
)

func bestClosingTime(customers string) int {
	penalty := 0
	n := len(customers)

	// Initial penalty if shop closes at hour 0:
	// all 'Y' customers will be missed
	for i := 0; i < n; i++ {
		if customers[i] == 'Y' {
			penalty++
		}
	}

	minPenalty := penalty
	bestHour := 0

	// Try closing at hour 1 to n
	for i := 0; i < n; i++ {
		if customers[i] == 'Y' {
			// One fewer missed customer
			penalty--
		} else {
			// One extra hour open with no customers
			penalty++
		}

		if penalty < minPenalty {
			minPenalty = penalty
			bestHour = i + 1
		}
	}

	return bestHour
}

func main() {
	// Test cases
	fmt.Println(bestClosingTime("YYNY"))   // Output: 2
	fmt.Println(bestClosingTime("NNNNN"))  // Output: 0
	fmt.Println(bestClosingTime("YYYY"))   // Output: 4
	fmt.Println(bestClosingTime("YNYNNY")) // Extra test
}
