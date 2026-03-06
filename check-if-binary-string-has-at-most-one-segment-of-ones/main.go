package main

import "fmt"

func checkOnesSegment(s string) bool {
	seenZero := false

	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			seenZero = true
		}

		if s[i] == '1' && seenZero {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(checkOnesSegment("1001")) // false
	fmt.Println(checkOnesSegment("110"))  // true
}
