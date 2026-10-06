package main

import "fmt"

func minAddToMakeValid(s string) int {
	open := 0
	add := 0

	for _, ch := range s {
		if ch == '(' {
			open++
		} else {
			if open > 0 {
				open--
			} else {
				add++
			}
		}
	}

	return add + open
}

func main() {
	s := "())"

	result := minAddToMakeValid(s)

	fmt.Println(result)
}
