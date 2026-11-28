package main

import (
	"fmt"
	"slices"
)

func maximumSaleItems(items [][]int, budget int) int {
	slices.SortFunc(items, func(a, b []int) int {
		return a[1] - b[1]
	})

	mx := 0
	for _, item := range items {
		mx = max(mx, item[0])
	}

	freq := make([]int, mx+1)

	for _, item := range items {
		freq[item[0]]++
	}

	// freq[i] = number of items whose factor is divisible by i
	for i := 1; i <= mx; i++ {
		if freq[i] > 0 {
			for j := i * 2; j <= mx; j += i {
				freq[i] += freq[j]
			}
		}
	}

	// Cheapest item
	mn := items[0][1]

	ans := 0

	for _, item := range items {
		factor, price := item[0], item[1]

		// If this item costs more than twice the cheapest item,
		// buying it cannot improve the answer.
		if price > mn*2 {
			break
		}

		// One purchase of this factor gives:
		// - the purchased item
		// - one free copy of every item whose factor is divisible by it
		cnt := freq[factor] - 1

		// Buy as many copies as possible with the remaining budget.
		take := min(cnt, budget/price)

		ans += take * 2
		budget -= take * price
	}

	// Use whatever budget remains to buy the cheapest item.
	ans += budget / mn

	return ans
}

func main() {
	items := [][]int{
		{2, 3},
		{4, 2},
		{6, 5},
	}

	budget := 6

	fmt.Println(maximumSaleItems(items, budget))
}