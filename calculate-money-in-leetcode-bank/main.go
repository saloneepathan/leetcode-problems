package main

import "fmt"

func totalMoney(n int) int {
	weeks := n / 7
	days := n % 7
	total := 0

	// Total for full weeks
	// Sum for all full weeks = 28*weeks + 7*(weeks-1)*weeks/2
	total += 28*weeks + 7*(weeks-1)*weeks/2

	// Remaining days of the next week
	start := weeks + 1
	for i := 0; i < days; i++ {
		total += start + i
	}

	return total
}

func main() {
	fmt.Println(totalMoney(4))  // Output: 10
	fmt.Println(totalMoney(10)) // Output: 37
	fmt.Println(totalMoney(20)) // Output: 96
}
