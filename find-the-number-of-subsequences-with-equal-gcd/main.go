package main

import (
	"fmt"
)

const (
	MOD = 1000000007
	MAX = 200
)

// Euclidean Algorithm
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func subsequencePairCount(nums []int) int {
	// dp[g1][g2] = number of ways
	// g1 = gcd of seq1
	// g2 = gcd of seq2
	dp := make([][]int, MAX+1)
	for i := range dp {
		dp[i] = make([]int, MAX+1)
	}

	// Initially both subsequences are empty
	dp[0][0] = 1

	for _, x := range nums {

		// Create next DP
		ndp := make([][]int, MAX+1)
		for i := range ndp {
			ndp[i] = make([]int, MAX+1)
			copy(ndp[i], dp[i]) // Option 1: Ignore current element
		}

		for g1 := 0; g1 <= MAX; g1++ {
			for g2 := 0; g2 <= MAX; g2++ {

				if dp[g1][g2] == 0 {
					continue
				}

				ways := dp[g1][g2]

				// Option 2: Put current element into seq1
				newG1 := x
				if g1 != 0 {
					newG1 = gcd(g1, x)
				}

				ndp[newG1][g2] = (ndp[newG1][g2] + ways) % MOD

				// Option 3: Put current element into seq2
				newG2 := x
				if g2 != 0 {
					newG2 = gcd(g2, x)
				}

				ndp[g1][newG2] = (ndp[g1][newG2] + ways) % MOD
			}
		}

		dp = ndp
	}

	answer := 0

	for g := 1; g <= MAX; g++ {
		answer = (answer + dp[g][g]) % MOD
	}

	return answer
}

func main() {
	nums1 := []int{1, 2, 3, 4}
	fmt.Println(subsequencePairCount(nums1)) // 10

	nums2 := []int{10, 20, 30}
	fmt.Println(subsequencePairCount(nums2)) // 2

	nums3 := []int{1, 1, 1, 1}
	fmt.Println(subsequencePairCount(nums3)) // 50
}
