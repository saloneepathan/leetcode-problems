package main

import "fmt"

func repeatedNTimes(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
		if i+2 < len(nums) && nums[i] == nums[i+2] {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}

func main() {
	fmt.Println(repeatedNTimes([]int{1, 2, 3, 3}))             // 3
	fmt.Println(repeatedNTimes([]int{2, 1, 2, 5, 3, 2}))       // 2
	fmt.Println(repeatedNTimes([]int{5, 1, 5, 2, 5, 3, 5, 4})) // 5
}
