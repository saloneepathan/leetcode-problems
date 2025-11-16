package main

import (
	"bufio"
	"fmt"
	"os"
)

func numSub(s string) int {
	const mod = 1_000_000_007
	res := 0
	streak := 0

	for i := 0; i < len(s); i++ {
		// Convert '0' -> 0 and '1' -> 1
		bit := int(s[i] - '0')

		// If bit = 1 → streak = streak + 1
		// If bit = 0 → streak = 0
		streak = (streak + bit) * bit

		res += streak
		if res >= mod {
			res -= mod
		}
	}

	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var s string
	fmt.Fscan(in, &s)
	fmt.Println(numSub(s))
}
