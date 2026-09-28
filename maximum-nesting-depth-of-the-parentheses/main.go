package main

import "fmt"

func maxDepth(s string) int {
	depth := 0
	maxDepth := 0

	for _, ch := range s {
		if ch == '(' {
			depth++
			if depth > maxDepth {
				maxDepth = depth
			}
		} else if ch == ')' {
			depth--
		}
	}

	return maxDepth
}

func main() {
	fmt.Println(maxDepth("(1+(2*3)+((8)/4))+1")) // 3
	fmt.Println(maxDepth("(1)+((2))+(((3)))"))   // 3
	fmt.Println(maxDepth("()(())((()()))"))      // 3
}
