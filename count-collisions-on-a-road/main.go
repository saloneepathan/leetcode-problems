package main

import (
	"fmt"
)

func countCollisions(directions string) int {
	n := len(directions)
	if n == 0 {
		return 0
	}

	// Skip leading L's — they move left forever without collision
	left := 0
	for left < n && directions[left] == 'L' {
		left++
	}

	// Skip trailing R's — they move right forever without collision
	right := n - 1
	for right >= 0 && directions[right] == 'R' {
		right--
	}

	collisions := 0

	// Inside the remaining segment:
	// Every R or L will eventually collide and become stationary.
	for i := left; i <= right; i++ {
		if directions[i] != 'S' {
			collisions++
		}
	}

	return collisions
}

func main() {
	fmt.Println(countCollisions("RLRSLL")) // Output: 5
	fmt.Println(countCollisions("LLRR"))   // Output: 0
}
