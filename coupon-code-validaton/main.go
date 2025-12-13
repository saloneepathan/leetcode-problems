package main

import (
	"fmt"
	"sort"
)

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	// Required business line order
	order := []string{"electronics", "grocery", "pharmacy", "restaurant"}

	// Valid business lines set
	validBusiness := map[string]bool{
		"electronics": true,
		"grocery":     true,
		"pharmacy":    true,
		"restaurant":  true,
	}

	// Group valid coupons by business line
	grouped := make(map[string][]string)

	for i := 0; i < len(code); i++ {
		// Check active
		if !isActive[i] {
			continue
		}

		// Check business line validity
		if !validBusiness[businessLine[i]] {
			continue
		}

		// Check code validity
		if !isValidCode(code[i]) {
			continue
		}

		grouped[businessLine[i]] = append(grouped[businessLine[i]], code[i])
	}

	// Build result following required order
	result := []string{}
	for _, bl := range order {
		if codes, ok := grouped[bl]; ok {
			sort.Strings(codes)
			result = append(result, codes...)
		}
	}

	return result
}

// Helper function to validate coupon code
func isValidCode(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, ch := range s {
		if !(ch >= 'a' && ch <= 'z' ||
			ch >= 'A' && ch <= 'Z' ||
			ch >= '0' && ch <= '9' ||
			ch == '_') {
			return false
		}
	}
	return true
}

func main() {
	// Example 1
	code1 := []string{"SAVE20", "", "PHARMA5", "SAVE@20"}
	businessLine1 := []string{"restaurant", "grocery", "pharmacy", "restaurant"}
	isActive1 := []bool{true, true, true, true}

	fmt.Println(validateCoupons(code1, businessLine1, isActive1))
	// Expected Output: [PHARMA5 SAVE20]

	// Example 2
	code2 := []string{"GROCERY15", "ELECTRONICS_50", "DISCOUNT10"}
	businessLine2 := []string{"grocery", "electronics", "invalid"}
	isActive2 := []bool{false, true, true}

	fmt.Println(validateCoupons(code2, businessLine2, isActive2))
	// Expected Output: [ELECTRONICS_50]
}
