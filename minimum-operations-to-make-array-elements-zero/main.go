package main

import (
	"fmt"
)

// prefixSum computes F(n) = sum of steps(x) for x in [1..n].
// steps(x) is the number of times x must be divided by 4 (floor) to become 0.
func prefixSum(n int64) int64 {
	if n <= 0 {
		return 0
	}

	var sum int64 = 0
	var k int64 = 1
	var base int64 = 1 // 4^(k-1)

	for {
		// range for step k: [4^(k-1), 4^k - 1]
		l := base
		r := base*4 - 1

		if n < l {
			break
		}

		// overlap with [1..n]
		if n >= r {
			count := r - l + 1
			sum += count * k
		} else {
			count := n - l + 1
			sum += count * k
			break
		}

		base *= 4
		k++
	}

	return sum
}

// Accept [][]int (common) and convert elements to int64 for math.
func minOperations(queries [][]int) int64 {
	var total int64 = 0

	for _, q := range queries {
		l := int64(q[0])
		r := int64(q[1])
		S := prefixSum(r) - prefixSum(l-1)
		ops := (S + 1) / 2 // ceil(S/2)
		total += ops
	}

	return total
}

func main() {
	// Example 1
	queries1 := [][]int{{1, 2}, {2, 4}}
	fmt.Println(minOperations(queries1)) // 3

	// Example 2
	queries2 := [][]int{{2, 6}}
	fmt.Println(minOperations(queries2)) // 4
}
