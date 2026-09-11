package main

import "fmt"

func totalNumbers(digits []int) int {
	// Count how many times each digit appears.
	count := make([]int, 10)

	for _, digit := range digits {
		count[digit]++
	}

	ans := 0

	// Try every 3-digit even number.
	for num := 100; num <= 998; num += 2 {
		a := num / 100       // Hundreds digit
		b := (num / 10) % 10 // Tens digit
		c := num % 10        // Ones digit

		// Check whether we have enough copies of each digit.
		used := make([]int, 10)
		used[a]++
		used[b]++
		used[c]++

		valid := true

		for d := 0; d <= 9; d++ {
			if used[d] > count[d] {
				valid = false
				break
			}
		}

		if valid {
			ans++
		}
	}

	return ans
}

func main() {
	fmt.Println(totalNumbers([]int{1, 2, 3, 4})) // 12
	fmt.Println(totalNumbers([]int{0, 2, 2}))    // 2
	fmt.Println(totalNumbers([]int{6, 6, 6}))    // 1
	fmt.Println(totalNumbers([]int{1, 3, 5}))    // 0
}
