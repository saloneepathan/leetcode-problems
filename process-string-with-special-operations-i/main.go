package main

import (
	"fmt"
)

func processStr(s string) string {
	result := ""

	for _, ch := range s {
		switch ch {
		case '*':
			if len(result) > 0 {
				result = result[:len(result)-1]
			}

		case '#':
			result += result

		case '%':
			runes := []rune(result)
			for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
				runes[i], runes[j] = runes[j], runes[i]
			}
			result = string(runes)

		default: // lowercase letter
			result += string(ch)
		}
	}

	return result
}

func main() {
	fmt.Println(processStr("a#b%*")) // ba
	fmt.Println(processStr("z*#"))   // ""
}
