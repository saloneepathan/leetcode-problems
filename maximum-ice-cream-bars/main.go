package main

import "fmt"

func maxIceCream(costs []int, coins int) int {
	// Find the maximum cost to create the counting array.
	maxCost := 0
	for _, cost := range costs {
		if cost > maxCost {
			maxCost = cost
		}
	}

	// count[i] = number of ice cream bars costing i coins.
	count := make([]int, maxCost+1)
	for _, cost := range costs {
		count[cost]++
	}

	bars := 0

	// Buy cheapest bars first.
	for cost := 1; cost <= maxCost; cost++ {
		if count[cost] == 0 {
			continue
		}

		// Number of bars we can afford at this price.
		canBuy := coins / cost

		if canBuy >= count[cost] {
			bars += count[cost]
			coins -= count[cost] * cost
		} else {
			bars += canBuy
			break
		}
	}

	return bars
}

func main() {
	fmt.Println(maxIceCream([]int{1, 3, 2, 4, 1}, 7))     // 4
	fmt.Println(maxIceCream([]int{10, 6, 8, 7, 7, 8}, 5)) // 0
	fmt.Println(maxIceCream([]int{1, 6, 3, 1, 2, 5}, 20)) // 6
}
