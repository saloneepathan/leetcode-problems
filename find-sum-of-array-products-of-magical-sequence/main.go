package main

import (
	"fmt"
	"math/big"
)

func magicalSum(m int, k int, nums []int) int {
	const MOD = 1e9 + 7
	n := len(nums)

	// dp[pos][sum] = sum of products where:
	// - We've filled pos positions
	// - The sum 2^seq[0] + 2^seq[1] + ... equals 'sum'
	// Note: sum can be very large, but we only care about set bit count

	// Since we can have up to m=30 identical indices, the max sum is
	// 30 * 2^49 which is huge. Instead, track (pos, setbitcount)
	// But we need to know which specific bits are set to correctly add new powers

	// Better: use big integers or track actual bit values
	// Use map with big.Int or string representation

	memo := make(map[string]int64)

	keyFunc := func(pos int, sum *big.Int) string {
		return fmt.Sprintf("%d:%s", pos, sum.String())
	}

	var solve func(int, *big.Int) int64
	solve = func(pos int, sum *big.Int) int64 {
		// Count set bits in sum
		bits := sum.BitLen()
		if sum.Sign() > 0 {
			bits = 0
			for i := 0; i < sum.BitLen(); i++ {
				if sum.Bit(i) == 1 {
					bits++
				}
			}
		}

		// Base case
		if pos == m {
			if bits == k {
				return 1
			}
			return 0
		}

		// Pruning
		if bits > k {
			return 0
		}
		if bits+m-pos < k {
			return 0
		}

		// Memo check
		memoKey := keyFunc(pos, sum)
		if val, exists := memo[memoKey]; exists {
			return val
		}

		result := int64(0)

		// Try each index
		for i := 0; i < n; i++ {
			// Add 2^i to sum
			newSum := new(big.Int).Set(sum)
			power := big.NewInt(1)
			power.Lsh(power, uint(i))
			newSum.Add(newSum, power)

			ways := solve(pos+1, newSum)
			contribution := (int64(nums[i]) * ways) % MOD
			result = (result + contribution) % MOD
		}

		memo[memoKey] = result
		return result
	}

	initialSum := big.NewInt(0)
	return int(solve(0, initialSum) % MOD)
}

func main() {
	fmt.Println(magicalSum(3, 2, []int{33}))                         // 35937
	fmt.Println(magicalSum(5, 5, []int{1, 10, 100, 10000, 1000000})) // 991600007
	fmt.Println(magicalSum(2, 2, []int{5, 4, 3, 2, 1}))              // 170
	fmt.Println(magicalSum(1, 1, []int{28}))                         // 28
}
