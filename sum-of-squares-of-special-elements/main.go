
package main

import "fmt"

func sumOfSquares(nums []int) int {
	n := len(nums)
	sum := 0

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			sum += nums[i-1] * nums[i-1]
		}
	}

	return sum
}

func main() {
	nums1 := []int{1, 2, 3, 4}
	fmt.Println(sumOfSquares(nums1)) // Output: 21

	nums2 := []int{2, 7, 1, 19, 18, 3}
	fmt.Println(sumOfSquares(nums2)) // Output: 63
}
