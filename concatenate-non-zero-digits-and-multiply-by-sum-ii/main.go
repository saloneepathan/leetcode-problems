package main

import "fmt"

func sumAndMultiply(s string, queries [][]int) []int {
	const MOD int64 = 1_000_000_007

	n := len(s)

	// pow10[i] = 10^i % MOD
	pow10 := make([]int64, n+1)
	pow10[0] = 1
	for i := 1; i <= n; i++ {
		pow10[i] = (pow10[i-1] * 10) % MOD
	}

	// Prefix arrays
	// val[i] = concatenated value of all non-zero digits in s[:i]
	// sum[i] = sum of all non-zero digits in s[:i]
	// cnt[i] = number of non-zero digits in s[:i]
	val := make([]int64, n+1)
	sum := make([]int, n+1)
	cnt := make([]int, n+1)

	for i := 0; i < n; i++ {
		val[i+1] = val[i]
		sum[i+1] = sum[i]
		cnt[i+1] = cnt[i]

		d := int(s[i] - '0')
		if d != 0 {
			val[i+1] = (val[i]*10 + int64(d)) % MOD
			sum[i+1] += d
			cnt[i+1]++
		}
	}

	ans := make([]int, len(queries))

	for i, q := range queries {
		l, r := q[0], q[1]

		// Number of non-zero digits in the substring
		k := cnt[r+1] - cnt[l]

		// Concatenated value of non-zero digits in the substring
		x := val[r+1] - (val[l]*pow10[k])%MOD
		if x < 0 {
			x += MOD
		}

		// Sum of non-zero digits in the substring
		digitSum := sum[r+1] - sum[l]

		ans[i] = int((x * int64(digitSum)) % MOD)
	}

	return ans
}

func main() {
	// Example 1
	s1 := "10203004"
	q1 := [][]int{{0, 7}, {1, 3}, {4, 6}}
	fmt.Println(sumAndMultiply(s1, q1)) // [12340 4 9]

	// Example 2
	s2 := "1000"
	q2 := [][]int{{0, 3}, {1, 1}}
	fmt.Println(sumAndMultiply(s2, q2)) // [1 0]

	// Example 3
	s3 := "9876543210"
	q3 := [][]int{{0, 9}}
	fmt.Println(sumAndMultiply(s3, q3)) // [444444137]
}
