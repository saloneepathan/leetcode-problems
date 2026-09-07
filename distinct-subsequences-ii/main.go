package main

import (
	"bufio"
	"fmt"
	"os"
)

func distinctSubseqII(s string) int {
	const MOD int64 = 1_000_000_007

	// dp includes the empty subsequence.
	dp := int64(1)

	// last[c] stores the dp value before the previous
	// occurrence of character c.
	last := make([]int64, 26)

	for _, ch := range s {
		c := ch - 'a'

		newDP := (2*dp - last[c] + MOD) % MOD

		last[c] = dp
		dp = newDP
	}

	// Remove the empty subsequence.
	return int((dp - 1 + MOD) % MOD)
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscan(in, &s)

	result := distinctSubseqII(s)
	fmt.Println(result)
}
