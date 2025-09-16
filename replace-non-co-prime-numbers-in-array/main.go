package main

import (
	"fmt"
)

// Function to find GCD (Euclidean algorithm)
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Function to find LCM using GCD
func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func replaceNonCoprimes(nums []int) []int {
	stack := []int{}

	for _, num := range nums {
		// push current number
		stack = append(stack, num)

		// keep merging while top two are non-coprime
		for len(stack) > 1 {
			n := len(stack)
			a, b := stack[n-2], stack[n-1]

			g := gcd(a, b)
			if g == 1 { // coprime, stop merging
				break
			}

			// merge into LCM
			stack = stack[:n-2] // pop last two
			stack = append(stack, lcm(a, b))
		}
	}

	return stack
}

func main() {
	// Example 1
	nums1 := []int{6, 4, 3, 2, 7, 6, 2}
	fmt.Println(replaceNonCoprimes(nums1)) // Output: [12 7 6]

	// Example 2
	nums2 := []int{2, 2, 1, 1, 3, 3, 3}
	fmt.Println(replaceNonCoprimes(nums2)) // Output: [2 1 1 3]
}
