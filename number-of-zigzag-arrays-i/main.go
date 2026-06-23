package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD int64 = 1_000_000_007

func zigZagArrays(n int, l int, r int) int {
	// Only the number of available values matters.
	m := r - l + 1

	// up[x]   = valid arrays ending at x where the last move was increasing
	// down[x] = valid arrays ending at x where the last move was decreasing
	up := make([]int64, m)
	down := make([]int64, m)

	// Base case: arrays of length 2
	for x := 0; x < m; x++ {
		up[x] = int64(x)           // values smaller than x
		down[x] = int64(m - 1 - x) // values greater than x
	}

	// Build arrays from length 3 to n
	for length := 3; length <= n; length++ {
		nextUp := make([]int64, m)
		nextDown := make([]int64, m)

		// nextUp[x] = sum of down[0...x-1]
		var prefix int64 = 0
		for x := 0; x < m; x++ {
			nextUp[x] = prefix
			prefix = (prefix + down[x]) % MOD
		}

		// nextDown[x] = sum of up[x+1...m-1]
		var suffix int64 = 0
		for x := m - 1; x >= 0; x-- {
			nextDown[x] = suffix
			suffix = (suffix + up[x]) % MOD
		}

		up = nextUp
		down = nextDown
	}

	var ans int64 = 0
	for x := 0; x < m; x++ {
		ans = (ans + up[x] + down[x]) % MOD
	}

	return int(ans)
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, l, r int
	fmt.Fscan(in, &n, &l, &r)

	fmt.Println(zigZagArrays(n, l, r))
}
