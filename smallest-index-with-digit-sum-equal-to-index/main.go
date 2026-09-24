package main

import "fmt"

func smallestIndex(nums []int) int {
	for i, num := range nums {
		sum := 0

		for num > 0 {
			sum += num % 10
			num /= 10
		}

		if sum == i {
			return i
		}
	}

	return -1
}

func main() {
	nums := []int{1, 3, 2}

	fmt.Println(smallestIndex(nums)) // 2
}