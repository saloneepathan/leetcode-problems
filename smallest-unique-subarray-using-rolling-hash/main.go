package main

import "fmt"

func smallestUniqueSubarray(nums []int) int {
	const MOD int64 = 1000000007
	const BASE int64 = 911382323

	n := len(nums)

	hash := make([]int64, n+1)
	powers := make([]int64, n+1)
	powers[0] = 1

	// Compute prefix hashes and powers of BASE
	for i, num := range nums {
		hash[i+1] = (hash[i]*BASE + int64(num)) % MOD
		powers[i+1] = powers[i] * BASE % MOD
	}

	// Check whether there is a unique subarray of given length.
	check := func(length int) bool {
		counts := make(map[int64]int, n-length+1)

		for left := 0; left+length <= n; left++ {
			right := left + length

			// Hash of nums[left:right]
			key := (hash[right] -
				hash[left]*powers[length]%MOD +
				MOD) % MOD

			counts[key]++
		}

		// At least one subarray occurs exactly once.
		for _, times := range counts {
			if times == 1 {
				return true
			}
		}

		return false
	}

	// Binary search for the smallest length
	// for which a unique subarray exists.
	lo, hi := 1, n

	for lo < hi {
		mid := lo + (hi-lo)/2

		if check(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return lo
}

func main() {
	nums := []int{1, 2, 1, 2, 3}

	answer := smallestUniqueSubarray(nums)

	fmt.Println(answer)
}