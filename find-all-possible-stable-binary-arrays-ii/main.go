package main

import (
	"fmt"
)

const MOD = 1000000007

func numberOfStableArrays(zero int, one int, limit int) int {
	memo := make([][][]int, zero+1)
	for i := range memo {
		memo[i] = make([][]int, one+1)
		for j := range memo[i] {
			memo[i][j] = []int{-1, -1}
		}
	}

	var dp func(int, int, int) int
	dp = func(z, o, lastBit int) int {
		if z == 0 {
			if lastBit == 0 || o > limit {
				return 0
			}
			return 1
		}

		if o == 0 {
			if lastBit == 1 || z > limit {
				return 0
			}
			return 1
		}

		if memo[z][o][lastBit] != -1 {
			return memo[z][o][lastBit]
		}

		res := 0

		if lastBit == 0 {
			res = (dp(z-1, o, 0) + dp(z-1, o, 1)) % MOD
			if z > limit {
				res = (res - dp(z-limit-1, o, 1) + MOD) % MOD
			}
		} else {
			res = (dp(z, o-1, 0) + dp(z, o-1, 1)) % MOD
			if o > limit {
				res = (res - dp(z, o-limit-1, 0) + MOD) % MOD
			}
		}

		memo[z][o][lastBit] = res
		return res
	}

	return (dp(zero, one, 0) + dp(zero, one, 1)) % MOD
}

func main() {
	zero := 2
	one := 1
	limit := 2

	result := numberOfStableArrays(zero, one, limit)
	fmt.Println("Number of Stable Binary Arrays:", result)
}
