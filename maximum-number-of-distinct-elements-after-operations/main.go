package main

import (
	"fmt"
	"sort"
)

func maxDistinctElements(nums []int, k int) int {
	sort.Ints(nums)

	last := int64(-1 << 60) // effectively negative infinity
	count := 0

	for _, num := range nums {
		low := int64(num) - int64(k)
		high := int64(num) + int64(k)

		if low > last {
			// we can safely place at low
			last = low
			count++
		} else if last < high {
			// shift just enough to get a new distinct number
			last++
			count++
		}
		// else: all possible positions are already taken, skip
	}

	return count
}

func main() {
	fmt.Println(maxDistinctElements([]int{1, 2, 2, 3, 3, 4}, 2))    // 6
	fmt.Println(maxDistinctElements([]int{4, 4, 4, 4}, 1))          // 3
	fmt.Println(maxDistinctElements([]int{10, 10, 10, 10}, 500000)) // should finish instantly
}
