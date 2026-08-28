package main

import (
	"fmt"
	"strings"
)

func lexPalindromicPermutation(s string, target string) string {
	n := len(s)

	// Count frequency of each character.
	cnt := make([]int, 26)
	for _, c := range s {
		cnt[c-'a']++
	}

	// A palindrome can have at most one character with odd frequency.
	oddChar := ""
	for i := 0; i < 26; i++ {
		if cnt[i]%2 == 1 {
			if oddChar != "" {
				return ""
			}
			oddChar = string(byte('a' + i))
		}

		// Only half of each pair is needed for the left half.
		cnt[i] /= 2
	}

	halfLen := n / 2

	// Build a palindrome from its left half.
	buildPalindrome := func(left string) string {
		right := reverseString(left)
		return left + oddChar + right
	}

	// Try to construct the answer greedily.
	//
	// At every position:
	// - Try characters from 'a' to 'z'.
	// - Temporarily place the character.
	// - Complete the remaining half in ascending order.
	// - If the resulting palindrome is > target, this is the
	//   smallest possible answer for this position.
	prefix := ""

	for i := 0; i < halfLen; i++ {
		found := false

		for j := 0; j < 26; j++ {
			if cnt[j] == 0 {
				continue
			}

			// Use this character.
			cnt[j]--

			// Construct the smallest possible completion.
			remaining := prefix + string(byte('a'+j))

			for k := 0; k < 26; k++ {
				if cnt[k] > 0 {
					remaining += strings.Repeat(
						string(byte('a'+k)),
						cnt[k],
					)
				}
			}

			candidate := buildPalindrome(remaining)

			if candidate > target {
				// Since j is tried from smallest to largest,
				// this is the smallest possible answer for
				// the current prefix.
				prefix += string(byte('a' + j))
				found = true
				break
			}

			// Undo the choice.
			cnt[j]++
		}

		if found {
			// Once the complete candidate is already greater,
			// we can greedily fill the rest with the smallest
			// available characters.
			for k := 0; k < 26; k++ {
				if cnt[k] > 0 {
					prefix += strings.Repeat(
						string(byte('a'+k)),
						cnt[k],
					)
					cnt[k] = 0
				}
			}

			return buildPalindrome(prefix)
		}
	}

	// If we reach here, even the largest possible palindrome
	// is not strictly greater than target.
	return ""
}

func reverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	tests := []struct {
		s      string
		target string
		want   string
	}{
		{
			s:      "baba",
			target: "abba",
			want:   "baab",
		},
		{
			s:      "baba",
			target: "bbaa",
			want:   "",
		},
		{
			s:      "abc",
			target: "abb",
			want:   "",
		},
		{
			s:      "aac",
			target: "abb",
			want:   "aca",
		},
		{
			s:      "aaaa",
			target: "aaaa",
			want:   "",
		},
		{
			s:      "aaaa",
			target: "aaab",
			want:   "",
		},
		{
			s:      "abba",
			target: "aaaa",
			want:   "abba",
		},
	}

	for _, test := range tests {
		got := lexPalindromicPermutation(test.s, test.target)

		fmt.Printf(
			"s=%q, target=%q -> got=%q, want=%q\n",
			test.s,
			test.target,
			got,
			test.want,
		)
	}
}