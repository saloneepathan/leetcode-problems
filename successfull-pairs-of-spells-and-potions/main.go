package main

import (
	"fmt"
	"sort"
)

// successfulPairs returns, for each spell, the number of potions that form a successful pair
// with that spell. Signature matches common Go LeetCode solutions:
// spells, potions are int slices and success is int64.
func successfulPairs(spells []int, potions []int, success int64) []int {
	n := len(potions)
	// Convert potions to int64 slice for safe arithmetic and sort it.
	p64 := make([]int64, n)
	for i := 0; i < n; i++ {
		p64[i] = int64(potions[i])
	}
	sort.Slice(p64, func(i, j int) bool { return p64[i] < p64[j] })

	res := make([]int, len(spells))
	for i, sp := range spells {
		s := int64(sp)
		if s == 0 {
			// if spell strength is 0, product will be 0 -> never reach success > 0
			res[i] = 0
			continue
		}
		// threshold = ceil(success / s)
		// Using integer math: (success + s - 1) / s
		threshold := (success + s - 1) / s

		// find first index in p64 with value >= threshold
		idx := sort.Search(n, func(j int) bool {
			return p64[j] >= threshold
		})
		res[i] = n - idx
	}
	return res
}

func main() {
	// Example from LeetCode:
	spells := []int{5, 1, 3}
	potions := []int{1, 2, 3, 4, 5}
	success := int64(7)
	fmt.Println(successfulPairs(spells, potions, success)) // expected [4,0,3]
}
