package main

import "fmt"

func hasSameDigits(s string) bool {
	for len(s) > 2 {
		next := make([]byte, len(s)-1)
		for i := 0; i < len(s)-1; i++ {
			sum := (int(s[i]-'0') + int(s[i+1]-'0')) % 10
			next[i] = byte(sum) + '0'
		}
		s = string(next)
	}
	return s[0] == s[1]
}

func main() {
	fmt.Println(hasSameDigits("3902"))  // true
	fmt.Println(hasSameDigits("34789")) // false
}
