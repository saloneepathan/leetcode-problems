package main

import "fmt"

func numWaterBottles(numBottles int, numExchange int) int {
	total := numBottles
	empty := numBottles

	for empty >= numExchange {
		// Exchange empty bottles for full ones
		newBottles := empty / numExchange
		total += newBottles
		// Update empty bottles count
		empty = empty%numExchange + newBottles
	}

	return total
}

func main() {
	fmt.Println(numWaterBottles(9, 3))  // Output: 13
	fmt.Println(numWaterBottles(15, 4)) // Output: 19
	fmt.Println(numWaterBottles(5, 5))  // Output: 6
	fmt.Println(numWaterBottles(2, 3))  // Output: 2
}
