package main

import (
	"fmt"
	"sort"
)

func asteroidsDestroyed(mass int, asteroids []int) bool {
	sort.Ints(asteroids)

	cur := int64(mass)
	for _, a := range asteroids {
		if cur < int64(a) {
			return false
		}
		cur += int64(a)
	}

	return true
}

func main() {
	fmt.Println(asteroidsDestroyed(10, []int{3, 9, 19, 5, 21})) // true
	fmt.Println(asteroidsDestroyed(5, []int{4, 9, 23, 4}))      // false
}
