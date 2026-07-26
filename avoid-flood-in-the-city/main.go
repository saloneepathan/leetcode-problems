package main

import (
	"fmt"
	"sort"
)

func avoidFlood(rains []int) []int {
	n := len(rains)
	ans := make([]int, n)
	lastRain := make(map[int]int) // lake -> last day it rained
	var dryDays []int             // indices of days when we can dry

	for i := 0; i < n; i++ {
		if rains[i] > 0 {
			lake := rains[i]
			ans[i] = -1

			if lastDay, exists := lastRain[lake]; exists {
				// Find the first dry day > lastDay
				idx := sort.Search(len(dryDays), func(j int) bool {
					return dryDays[j] > lastDay
				})
				if idx == len(dryDays) {
					return []int{} // no suitable dry day, flood unavoidable
				}

				// Use this dry day to dry 'lake'
				ans[dryDays[idx]] = lake
				dryDays = append(dryDays[:idx], dryDays[idx+1:]...)
			}

			lastRain[lake] = i
		} else {
			// Mark this as a day available for drying
			dryDays = append(dryDays, i)
			ans[i] = 1 // placeholder, may be updated later
		}
	}

	return ans
}

func main() {
	fmt.Println(avoidFlood([]int{1, 2, 3, 4}))       // [-1 -1 -1 -1]
	fmt.Println(avoidFlood([]int{1, 2, 0, 0, 2, 1})) // [-1 -1 2 1 -1 -1]
	fmt.Println(avoidFlood([]int{1, 2, 0, 1, 2}))    // []
}
