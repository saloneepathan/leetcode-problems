package main

import (
	"fmt"
	"math"
)

func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {

	can := func(T int64) bool {
		var total int64 = 0

		for _, t := range workerTimes {

			k := float64(2*T) / float64(t)

			x := int64((math.Sqrt(1+4*k) - 1) / 2)

			total += x

			if total >= int64(mountainHeight) {
				return true
			}
		}

		return false
	}

	left, right := int64(0), int64(1e18)
	ans := right

	for left <= right {
		mid := (left + right) / 2

		if can(mid) {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}

func main() {
	fmt.Println(minNumberOfSeconds(4, []int{2, 1, 1}))     // 3
	fmt.Println(minNumberOfSeconds(10, []int{3, 2, 2, 4})) // 12
	fmt.Println(minNumberOfSeconds(5, []int{1}))           // 15
}
