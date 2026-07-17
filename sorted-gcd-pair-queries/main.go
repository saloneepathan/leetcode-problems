package main

import (
	"fmt"
	"sort"
)

func gcdValues(nums []int, queries []int64) []int {
	maxVal := 0
	for _, x := range nums {
		if x > maxVal {
			maxVal = x
		}
	}

	freq := make([]int64, maxVal+1)
	for _, x := range nums {
		freq[x]++
	}

	// cnt[g] = how many numbers are divisible by g
	cnt := make([]int64, maxVal+1)

	for g := 1; g <= maxVal; g++ {
		for multiple := g; multiple <= maxVal; multiple += g {
			cnt[g] += freq[multiple]
		}
	}

	// exact[g] = number of pairs with gcd exactly g
	exact := make([]int64, maxVal+1)

	for g := maxVal; g >= 1; g-- {
		c := cnt[g]
		exact[g] = c * (c - 1) / 2

		for multiple := 2 * g; multiple <= maxVal; multiple += g {
			exact[g] -= exact[multiple]
		}
	}

	// Prefix sums of gcd frequencies.
	prefix := make([]int64, maxVal+1)
	for g := 1; g <= maxVal; g++ {
		prefix[g] = prefix[g-1] + exact[g]
	}

	ans := make([]int, len(queries))

	for i, q := range queries {
		// Find first gcd value with prefix > q.
		idx := sort.Search(maxVal+1, func(x int) bool {
			return prefix[x] > q
		})
		ans[i] = idx
	}

	return ans
}

func main() {
	fmt.Println(gcdValues([]int{2, 3, 4}, []int64{0, 2, 2}))
	// [1 2 2]

	fmt.Println(gcdValues([]int{4, 4, 2, 1}, []int64{5, 3, 1, 0}))
	// [4 2 1 1]

	fmt.Println(gcdValues([]int{2, 2}, []int64{0, 0}))
	// [2 2]
}