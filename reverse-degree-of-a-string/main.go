package main

import "fmt"

func reverseDegree(s string) int {
	result := 0

	for i := 0; i < len(s); i++ {
		// 'a' = 26, 'b' = 25, ..., 'z' = 1
		reverseValue := int('z'-s[i]) + 1
		position := i + 1

		result += reverseValue * position
	}

	return result
}

func main() {
	fmt.Println(reverseDegree("abc"))  // 148
	fmt.Println(reverseDegree("zaza")) // 160
}