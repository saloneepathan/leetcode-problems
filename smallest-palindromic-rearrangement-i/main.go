package main

import "fmt"

func smallestPalindrome(s string) string {
	var freq [26]int
	for _, ch := range s {
		freq[ch-'a']++
	}

	// Build the left half of the palindrome.
	left := make([]byte, 0, len(s)/2)
	for i := 0; i < 26; i++ {
		for j := 0; j < freq[i]/2; j++ {
			left = append(left, byte('a'+i))
		}
	}

	// Find the middle character (if the length is odd).
	middle := byte(0)
	for i := 0; i < 26; i++ {
		if freq[i]%2 == 1 {
			middle = byte('a' + i)
			break
		}
	}

	// Construct the answer.
	ans := make([]byte, 0, len(s))
	ans = append(ans, left...)

	if middle != 0 {
		ans = append(ans, middle)
	}

	// Append the reverse of the left half.
	for i := len(left) - 1; i >= 0; i-- {
		ans = append(ans, left[i])
	}

	return string(ans)
}

func main() {
	testCases := []string{
		"z",
		"babab",
		"daccad",
		"aabbccbbaa",
	}

	for _, s := range testCases {
		fmt.Printf("Input: %s -> Output: %s\n", s, smallestPalindrome(s))
	}
}
