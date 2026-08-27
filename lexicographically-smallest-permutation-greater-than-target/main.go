package main

import "fmt"

func lexGreaterPermutation(s string, target string) string {
	n := len(s)

	// Frequency of characters in s.
	cnt := make([]int, 26)
	for _, ch := range s {
		cnt[ch-'a']++
	}

	// Try every position as the first position where
	// our answer differs from target.
	for i := 0; i < n; i++ {
		used := make([]int, 26)
		ok := true

		// The prefix [0, i) must equal target.
		for j := 0; j < i; j++ {
			c := int(target[j] - 'a')
			used[c]++

			if used[c] > cnt[c] {
				ok = false
				break
			}
		}

		// If target's prefix cannot be formed, no later position
		// can work either.
		if !ok {
			break
		}

		// Characters remaining after constructing target[:i].
		remaining := make([]int, 26)
		copy(remaining, cnt)

		for c := 0; c < 26; c++ {
			remaining[c] -= used[c]
		}

		// Choose the smallest character strictly greater than
		// target[i].
		targetChar := int(target[i] - 'a')

		for c := targetChar + 1; c < 26; c++ {
			if remaining[c] == 0 {
				continue
			}

			remaining[c]--

			// Build the smallest possible answer.
			ans := make([]byte, 0, n)

			// Equal prefix.
			ans = append(ans, target[:i]...)

			// First character that makes the result greater.
			ans = append(ans, byte('a'+c))

			// Smallest possible suffix.
			for x := 0; x < 26; x++ {
				for remaining[x] > 0 {
					ans = append(ans, byte('a'+x))
					remaining[x]--
				}
			}

			return string(ans)
		}
	}

	return ""
}

func main() {
	tests := []struct {
		s      string
		target string
	}{
		{"abc", "bba"},
		{"leet", "code"},
		{"baba", "bbaa"},
		{"abc", "abc"},
		{"a", "a"},
		{"ab", "aa"},
		{"aba", "aab"},
	}

	for _, test := range tests {
		result := lexGreaterPermutation(test.s, test.target)

		fmt.Printf(
			"s = %q, target = %q -> %q\n",
			test.s,
			test.target,
			result,
		)
	}
}