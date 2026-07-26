package main

import "fmt"

// TriangularSum computes the triangular sum of the array
func triangularSum(nums []int) int {
	n := len(nums)
	for n > 1 {
		for i := 0; i < n-1; i++ {
			nums[i] = (nums[i] + nums[i+1]) % 10
		}
		n--
	}
	return nums[0]
}

func main() {
	fmt.Println(triangularSum([]int{1, 2, 3, 4, 5})) // Output: 8
	fmt.Println(triangularSum([]int{5}))             // Output: 5
}
