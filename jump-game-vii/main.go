package main

import "fmt"

func canReach(s string, minJump int, maxJump int) bool {
	n := len(s)

	dp := make([]bool, n)
	dp[0] = true

	// count of reachable positions in current sliding window
	pre := 0

	for i := 1; i < n; i++ {
		// add new position entering the window
		if i-minJump >= 0 && dp[i-minJump] {
			pre++
		}

		// remove old position leaving the window
		if i-maxJump-1 >= 0 && dp[i-maxJump-1] {
			pre--
		}

		// current position is reachable
		if s[i] == '0' && pre > 0 {
			dp[i] = true
		}
	}

	return dp[n-1]
}

func main() {
	fmt.Println(canReach("011010", 2, 3))   // true
	fmt.Println(canReach("01101110", 2, 3)) // false
}
