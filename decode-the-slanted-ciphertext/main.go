package main

import (
	"fmt"
)

func decodeCiphertext(encodedText string, rows int) string {
	if rows == 1 {
		return encodedText
	}

	n := len(encodedText)
	if n == 0 {
		return ""
	}

	cols := n / rows

	// Read diagonally without building matrix
	result := make([]byte, 0, n)

	for startCol := 0; startCol < cols; startCol++ {
		r, c := 0, startCol
		for r < rows && c < cols {
			result = append(result, encodedText[r*cols+c])
			r++
			c++
		}
	}

	// Trim trailing spaces
	i := len(result) - 1
	for i >= 0 && result[i] == ' ' {
		i--
	}

	if i < 0 {
		return ""
	}

	return string(result[:i+1])
}

func main() {
	// Test cases
	fmt.Println(decodeCiphertext("ch   ie   pr", 3))             // cipher
	fmt.Println(decodeCiphertext("iveo    eed   l te   olc", 4)) // i love leetcode
	fmt.Println(decodeCiphertext("coding", 1))                   // coding
	fmt.Println(decodeCiphertext("", 5))                         // ""
}
