package main

import (
	"fmt"
)

func maxPower(stations []int, r int, k int) int64 {
	n := len(stations)

	// Prefix sum for quick range power calculation
	prefix := make([]int64, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + int64(stations[i])
	}

	// Compute initial power for each city
	power := make([]int64, n)
	for i := 0; i < n; i++ {
		left := i - r
		if left < 0 {
			left = 0
		}
		right := i + r
		if right >= n {
			right = n - 1
		}
		power[i] = prefix[right+1] - prefix[left]
	}

	// Binary search for the maximum achievable minimum power
	low, high := int64(0), prefix[n]+int64(k)
	var ans int64

	for low <= high {
		mid := (low + high) >> 1
		if canAchieve(mid, power, r, int64(k)) {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return ans
}

func canAchieve(target int64, power []int64, r int, k int64) bool {
	n := len(power)
	diff := make([]int64, n)
	var used int64

	for i := 0; i < n; i++ {
		if i > 0 {
			diff[i] += diff[i-1] // accumulate previous additions
		}
		cur := power[i] + diff[i]
		if cur < target {
			need := target - cur
			if need > k-used {
				return false
			}
			used += need
			diff[i] += need
			if end := i + 2*r + 1; end < n {
				diff[end] -= need
			}
		}
	}
	return true
}

func main() {
	// Example 1
	stations := []int{1, 2, 4, 5, 0}
	r := 1
	k := 2
	fmt.Println(maxPower(stations, r, k)) // Expected: 5

	// Example 2
	stations2 := []int{4, 4, 4, 4}
	r2 := 0
	k2 := 3
	fmt.Println(maxPower(stations2, r2, k2)) // Expected: 4
}
