package main

import (
	"fmt"
)

func numberOfBeams(bank []string) int {
	prevCount := 0
	totalBeams := 0

	for _, row := range bank {
		// Count number of '1's in the current row
		currCount := 0
		for _, ch := range row {
			if ch == '1' {
				currCount++
			}
		}

		// If the row has no devices, skip
		if currCount == 0 {
			continue
		}

		// Add beams between previous and current non-empty rows
		totalBeams += prevCount * currCount

		// Update previous row device count
		prevCount = currCount
	}

	return totalBeams
}

func main() {
	// Test cases
	bank1 := []string{"011001", "000000", "010100", "001000"}
	fmt.Println(numberOfBeams(bank1)) // Expected: 8

	bank2 := []string{"000", "111", "000"}
	fmt.Println(numberOfBeams(bank2)) // Expected: 0

	bank3 := []string{"1", "0", "1"}
	fmt.Println(numberOfBeams(bank3)) // Expected: 1

	bank4 := []string{"10", "01", "10", "01"}
	fmt.Println(numberOfBeams(bank4)) // Expected: 6
}
