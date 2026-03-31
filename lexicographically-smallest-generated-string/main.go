package main

import (
	"fmt"
)

func generateString(str1 string, str2 string) string {
	n, m := len(str1), len(str2)
	s := make([]byte, n+m-1)
	fixed := make([]int, n+m-1)

	// initialize with 'a'
	for i := range s {
		s[i] = 'a'
	}

	// Handle 'T' constraints
	for i := 0; i < n; i++ {
		if str1[i] == 'T' {
			for j := i; j < i+m; j++ {
				if fixed[j] == 1 && s[j] != str2[j-i] {
					return ""
				}
				s[j] = str2[j-i]
				fixed[j] = 1
			}
		}
	}

	// Handle 'F' constraints
	for i := 0; i < n; i++ {
		if str1[i] == 'F' {
			flag := false
			idx := -1

			for j := i + m - 1; j >= i; j-- {
				if str2[j-i] != s[j] {
					flag = true
				}
				if idx == -1 && fixed[j] == 0 {
					idx = j
				}
			}

			if flag {
				continue
			} else if idx != -1 {
				s[idx] = 'b'
			} else {
				return ""
			}
		}
	}

	return string(s)
}

func main() {
	// Test cases
	fmt.Println(generateString("TFT", "ab"))  // Example 1
	fmt.Println(generateString("TT", "abc"))  // Example 2
	fmt.Println(generateString("FFF", "aa"))  // Example 3
	fmt.Println(generateString("TFTF", "ab")) // Example 4
}
