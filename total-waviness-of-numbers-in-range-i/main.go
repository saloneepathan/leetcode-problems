package main

import (
	"fmt"
)

func totalWaviness(num1 int, num2 int) int {
	ans := 0

	for x := num1; x <= num2; x++ {
		ans += waviness(x)
	}

	return ans
}

func waviness(x int) int {
	digits := make([]int, 0)

	for x > 0 {
		digits = append(digits, x%10)
		x /= 10
	}

	// Reverse digits to normal order
	for l, r := 0, len(digits)-1; l < r; l, r = l+1, r-1 {
		digits[l], digits[r] = digits[r], digits[l]
	}

	if len(digits) < 3 {
		return 0
	}

	count := 0

	for i := 1; i < len(digits)-1; i++ {
		if (digits[i] > digits[i-1] && digits[i] > digits[i+1]) ||
			(digits[i] < digits[i-1] && digits[i] < digits[i+1]) {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(totalWaviness(120, 130))   // 3
	fmt.Println(totalWaviness(198, 202))   // 3
	fmt.Println(totalWaviness(4848, 4848)) // 2
}
