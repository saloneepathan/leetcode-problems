package main

import (
	"fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	// Result string
	var result string

	// Handle sign
	if (numerator < 0) != (denominator < 0) {
		result += "-"
	}

	// Convert to int64 to prevent overflow
	num := int64(numerator)
	den := int64(denominator)
	if num < 0 {
		num = -num
	}
	if den < 0 {
		den = -den
	}

	// Integer part
	integerPart := num / den
	result += strconv.FormatInt(integerPart, 10)

	// Remainder
	remainder := num % den
	if remainder == 0 {
		return result
	}

	result += "."

	// Map to store remainder and its position in the decimal part
	remainderMap := map[int64]int{}
	decimalPart := ""

	for remainder != 0 {
		// If remainder repeats
		if pos, exists := remainderMap[remainder]; exists {
			// Insert parentheses
			return result + decimalPart[:pos] + "(" + decimalPart[pos:] + ")"
		}
		// Record the position of this remainder
		remainderMap[remainder] = len(decimalPart)

		remainder *= 10
		decimalPart += strconv.FormatInt(remainder/den, 10)
		remainder %= den
	}

	return result + decimalPart
}

func main() {
	fmt.Println(fractionToDecimal(1, 2))   // "0.5"
	fmt.Println(fractionToDecimal(2, 1))   // "2"
	fmt.Println(fractionToDecimal(4, 333)) // "0.(012)"
	fmt.Println(fractionToDecimal(1, 6))   // "0.1(6)"
	fmt.Println(fractionToDecimal(-50, 8)) // "-6.25"
	fmt.Println(fractionToDecimal(7, -12)) // "-0.58(3)"
}
