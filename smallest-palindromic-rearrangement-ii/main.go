package main

import (
	"fmt"
)

func smallestPalindrome(s string, k int) string {
	partition := len(s) / 2
	bucket := make([]int, 26)

	// Since s is already a palindrome, counting the first half is enough.
	for i := 0; i < partition; i++ {
		bucket[s[i]-'a']++
	}

	// Computes nCm, stopping early if it exceeds kVal.
	comb := func(n, m, kVal int) int {
		res := 1

		if n-m < m {
			m = n - m
		}

		for i := 1; i <= m; i++ {
			res = res * (n - i + 1) / i
			if res > kVal {
				return kVal + 1
			}
		}
		return res
	}

	// Counts the number of distinct permutations possible.
	var permutations func(int) int
	permutations = func(rem int) int {
		ways := 1

		for i := 0; i < 26; i++ {
			if bucket[i] == 0 {
				continue
			}

			ways *= comb(rem, bucket[i], k)

			if ways > k {
				return k + 1
			}

			rem -= bucket[i]
		}

		return ways
	}

	left := make([]byte, 0, partition)
	startIndex := 1 // k is 1-indexed.

	for pos := 0; pos < partition; pos++ {
		found := false

		for i := 0; i < 26; i++ {
			if bucket[i] == 0 {
				continue
			}

			bucket[i]--

			ways := permutations(partition - pos - 1)

			// The kth palindrome lies in this range.
			if startIndex+ways > k {
				left = append(left, byte(i+'a'))
				found = true
				break
			}

			// Otherwise, skip these many permutations.
			startIndex += ways
			bucket[i]++
		}

		if !found {
			return ""
		}
	}

	totalLen := len(s)
	res := make([]byte, totalLen)

	for i := 0; i < partition; i++ {
		res[i] = left[i]
		res[totalLen-1-i] = left[i]
	}

	// Middle character remains unchanged.
	if totalLen%2 != 0 {
		res[partition] = s[partition]
	}

	return string(res)
}

func main() {
	tests := []struct {
		s string
		k int
	}{
		{"abba", 1},
		{"abba", 2},
		{"aaaa", 1},
		{"aabaa", 1},
		{"aabbccbbaa", 1},
		{"aabbccbbaa", 2},
	}

	for _, t := range tests {
		fmt.Printf("s = %q, k = %d -> %q\n",
			t.s, t.k, smallestPalindrome(t.s, t.k))
	}
}
