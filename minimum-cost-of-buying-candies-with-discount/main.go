package main

import (
	"fmt"
	"sort"
)

func minimumCost(cost []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(cost)))

	ans := 0
	for i := 0; i < len(cost); i++ {
		if i%3 != 2 { // every third candy is free
			ans += cost[i]
		}
	}

	return ans
}

func main() {
	fmt.Println(minimumCost([]int{1, 2, 3}))          // 5
	fmt.Println(minimumCost([]int{6, 5, 7, 9, 2, 2})) // 23
	fmt.Println(minimumCost([]int{5, 5}))             // 10
}
