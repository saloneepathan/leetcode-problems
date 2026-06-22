package main

import (
	"fmt"
)

func maxNumberOfBalloons(text string) int {
	// Array to store frequency of each lowercase letter.
	// index 0 = 'a', index 1 = 'b', ..., index 25 = 'z'
	count := make([]int, 26)

	// Count every character in text.
	for _, ch := range text {
		count[ch-'a']++
	}

	// Get how many of each required character we have.
	b := count['b'-'a']
	a := count['a'-'a']
	l := count['l'-'a'] / 2 // Need 2 l's per balloon
	o := count['o'-'a'] / 2 // Need 2 o's per balloon
	n := count['n'-'a']

	// The smallest available count decides the answer.
	answer := b

	if a < answer {
		answer = a
	}

	if l < answer {
		answer = l
	}

	if o < answer {
		answer = o
	}

	if n < answer {
		answer = n
	}

	return answer
}

func main() {
	text1 := "nlaebolko"
	fmt.Println(maxNumberOfBalloons(text1)) // Output: 1

	text2 := "loonbalxballpoon"
	fmt.Println(maxNumberOfBalloons(text2)) // Output: 2

	text3 := "leetcode"
	fmt.Println(maxNumberOfBalloons(text3)) // Output: 0
}
