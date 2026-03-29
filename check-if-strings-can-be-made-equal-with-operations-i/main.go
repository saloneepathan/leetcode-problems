package main

import (
	"fmt"
)

func canBeEqual(s1 string, s2 string) bool {
	evenMatch := (s1[0] == s2[0] && s1[2] == s2[2]) ||
		(s1[0] == s2[2] && s1[2] == s2[0])

	oddMatch := (s1[1] == s2[1] && s1[3] == s2[3]) ||
		(s1[1] == s2[3] && s1[3] == s2[1])

	return evenMatch && oddMatch
}

func main() {
	// Test cases
	fmt.Println(canBeEqual("abcd", "cdab")) // true
	fmt.Println(canBeEqual("abcd", "dacb")) // false
	fmt.Println(canBeEqual("abdc", "acbd")) // true
	fmt.Println(canBeEqual("aaaa", "aaaa")) // true
}
