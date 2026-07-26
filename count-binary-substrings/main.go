package main

import (
	"fmt"
)

func countBinarySubstrings(s string) int {

	prevGroup := 0
	currGroup := 1
	result := 0

	for i := 1; i < len(s); i++ {

		if s[i] == s[i-1] {
			currGroup++
		} else {

			if prevGroup < currGroup {
				result += prevGroup
			} else {
				result += currGroup
			}

			prevGroup = currGroup
			currGroup = 1
		}
	}

	if prevGroup < currGroup {
		result += prevGroup
	} else {
		result += currGroup
	}

	return result
}

func main() {

	s1 := "00110011"
	s2 := "10101"

	fmt.Println("Input:", s1)
	fmt.Println("Output:", countBinarySubstrings(s1))

	fmt.Println("Input:", s2)
	fmt.Println("Output:", countBinarySubstrings(s2))

}
