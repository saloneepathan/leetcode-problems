package main

import (
	"fmt"
)

func numOfWays(n int) int {
	const MOD int64 = 1_000_000_007

	// Base case: first row
	var aba int64 = 6 // patterns like A B A
	var abc int64 = 6 // patterns like A B C

	for i := 2; i <= n; i++ {
		newAba := (aba*3 + abc*2) % MOD
		newAbc := (aba*2 + abc*2) % MOD
		aba = newAba
		abc = newAbc
	}

	return int((aba + abc) % MOD)
}

func main() {
	fmt.Println(numOfWays(1))    // 12
	fmt.Println(numOfWays(5000)) // 30228214
}
