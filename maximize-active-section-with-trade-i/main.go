package main

import "fmt"

func maxActiveSectionsAfterTrade(s string) int {
	// Count original active sections.
	ones := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			ones++
		}
	}

	// Add virtual '1' at both ends.
	t := "1" + s + "1"

	// Store the previous two runs.
	var prev2Char, prev1Char byte
	prev2Len, prev1Len := 0, 0

	maxGain := 0

	for i := 0; i < len(t); {
		j := i
		for j < len(t) && t[j] == t[i] {
			j++
		}

		curChar := t[i]
		curLen := j - i

		// Pattern: 0 ... 1 ... 0
		if prev2Char == '0' && prev1Char == '1' && curChar == '0' {
			gain := prev2Len + curLen
			if gain > maxGain {
				maxGain = gain
			}
		}

		// Shift window.
		prev2Char, prev2Len = prev1Char, prev1Len
		prev1Char, prev1Len = curChar, curLen

		i = j
	}

	return ones + maxGain
}

func main() {
	testCases := []string{
		"01",
		"0100",
		"1000100",
		"01010",
		"1111",
		"0000",
	}

	for _, s := range testCases {
		fmt.Printf("Input: %-8s Output: %d\n", s, maxActiveSectionsAfterTrade(s))
	}
}
