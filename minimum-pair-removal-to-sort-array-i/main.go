package main

import "fmt"

// Check if array is non-decreasing
func isNonDecreasing(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return false
		}
	}
	return true
}

func minimumPairRemoval(nums []int) int {
	operations := 0

	for !isNonDecreasing(nums) {
		minSum := nums[0] + nums[1]
		index := 0

		// Find adjacent pair with minimum sum
		for i := 1; i < len(nums)-1; i++ {
			sum := nums[i] + nums[i+1]
			if sum < minSum {
				minSum = sum
				index = i
			}
		}

		// Replace the pair with their sum
		newNums := make([]int, 0, len(nums)-1)
		newNums = append(newNums, nums[:index]...)
		newNums = append(newNums, minSum)
		newNums = append(newNums, nums[index+2:]...)

		nums = newNums
		operations++
	}

	return operations
}

func main() {
	fmt.Println(minimumPairRemoval([]int{5, 2, 3, 1})) // Output: 2
	fmt.Println(minimumPairRemoval([]int{1, 2, 2}))    // Output: 0
}
