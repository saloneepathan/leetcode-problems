package main

import "fmt"

func grayCode(n int) []int {
	result := make([]int, 0, 1<<n)

	for i := 0; i < (1 << n); i++ {
		gray := i ^ (i >> 1)
		result = append(result, gray)
	}

	return result
}

func main() {
	n := 2

	result := grayCode(n)

	fmt.Println(result)
}
