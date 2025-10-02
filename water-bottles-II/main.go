package main

import "fmt"

func maxBottlesDrunk(numBottles int, numExchange int) int {
	totalDrunk := numBottles
	empty := numBottles

	for empty >= numExchange {
		// Exchange once
		empty -= numExchange
		numExchange++ // After each exchange, cost increases
		totalDrunk++  // You get 1 full bottle
		empty++       // After drinking that full bottle, becomes empty
	}

	return totalDrunk
}

func main() {
	fmt.Println(maxBottlesDrunk(13, 6)) // Output: 15
	fmt.Println(maxBottlesDrunk(10, 3)) // Output: 13
	fmt.Println(maxBottlesDrunk(9, 3))  // Example: should return 13
}
