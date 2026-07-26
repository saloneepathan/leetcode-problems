package main

import "fmt"

func sumFourDivisors(nums []int) int {
	totalSum := 0

	for _, n := range nums {
		divCount := 0
		divSum := 0

		for i := 1; i*i <= n; i++ {
			if n%i == 0 {
				d1 := i
				d2 := n / i

				if d1 == d2 {
					divCount++
					divSum += d1
				} else {
					divCount += 2
					divSum += d1 + d2
				}

				// Early break if more than 4 divisors
				if divCount > 4 {
					break
				}
			}
		}

		if divCount == 4 {
			totalSum += divSum
		}
	}

	return totalSum
}

func main() {
	fmt.Println(sumFourDivisors([]int{21, 4, 7}))      // Output: 32
	fmt.Println(sumFourDivisors([]int{21, 21}))        // Output: 64
	fmt.Println(sumFourDivisors([]int{1, 2, 3, 4, 5})) // Output: 0
}
