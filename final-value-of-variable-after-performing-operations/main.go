package main

import "fmt"

func finalValueAfterOperations(operations []string) int {
	x := 0
	for _, op := range operations {
		if op == "++X" || op == "X++" {
			x++
		} else {
			x--
		}
	}
	return x
}

func main() {
	fmt.Println(finalValueAfterOperations([]string{"--X", "X++", "X++"}))        // 1
	fmt.Println(finalValueAfterOperations([]string{"++X", "++X", "X++"}))        // 3
	fmt.Println(finalValueAfterOperations([]string{"X++", "++X", "--X", "X--"})) // 0
}
