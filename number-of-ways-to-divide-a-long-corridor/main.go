package main

import (
	"fmt"
)

func numberOfWays(corridor string) int {
	const MOD = 1000000007

	// Store indices of all seats
	seats := []int{}
	for i, ch := range corridor {
		if ch == 'S' {
			seats = append(seats, i)
		}
	}

	// Invalid if no seats or odd number of seats
	if len(seats) == 0 || len(seats)%2 != 0 {
		return 0
	}

	ways := 1

	// Multiply choices between consecutive seat-pairs
	for i := 2; i < len(seats); i += 2 {
		plantsBetween := seats[i] - seats[i-1] - 1
		ways = (ways * (plantsBetween + 1)) % MOD
	}

	return ways
}

func main() {
	// Test cases
	fmt.Println(numberOfWays("SSPPSPS"))  // Expected: 3
	fmt.Println(numberOfWays("PPSPSP"))   // Expected: 1
	fmt.Println(numberOfWays("S"))        // Expected: 0
	fmt.Println(numberOfWays("PPSSPPSS")) // Expected: 3
}
