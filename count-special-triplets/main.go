package main

import "fmt"

func specialTriplets(nums []int) int {
	const MOD = 1_000_000_007
	const MAX = 200000 // because nums[i] ≤ 100000 → 2 * nums[i] ≤ 200000

	left := make([]int, MAX+1)
	right := make([]int, MAX+1)

	// Count all values in right initially
	for _, v := range nums {
		right[v]++
	}

	result := 0

	for _, mid := range nums {
		// mid moves from right to middle
		right[mid]--

		target := mid * 2
		if target <= MAX {
			result = (result + left[target]*right[target]) % MOD
		}

		// mid moves into left
		left[mid]++
	}

	return result
}

// ---------- Test ----------
func main() {
	fmt.Println(specialTriplets([]int{6, 3, 6}))       // 1
	fmt.Println(specialTriplets([]int{0, 1, 0, 0}))    // 1
	fmt.Println(specialTriplets([]int{8, 4, 2, 8, 4})) // 2
}
