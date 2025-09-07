package main

import "fmt"

func sumZero(n int) []int {
	res := make([]int, 0, n)
	for i := 1; i <= n/2; i++ {
		res = append(res, -i, i)
	}
	if n%2 != 0 {
		res = append(res, 0)
	}
	return res
}

func main() {
	fmt.Println(sumZero(5)) // Example output: [-1 1 -2 2 0]
	fmt.Println(sumZero(3)) // Example output: [-1 1 0]
	fmt.Println(sumZero(1)) // Example output: [0]
}
